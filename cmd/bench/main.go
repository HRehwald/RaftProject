// Package main implements a benchmarking CLI for the raft-kv-store.
//
// This is NOT part of the MIT 6.5840 lab code — it's the piece that
// differentiates this project from a class assignment. It drives load
// against a running (or in-process simulated) cluster, optionally injects
// failures mid-run, and reports throughput/latency.
//
// Usage (once raft1/ and kvraft1/ are filled in from the MIT starter code
// and your own KV server implementation is written):
//
//	go run cmd/bench/main.go --cluster-size 3 --duration 30s
//	go run cmd/bench/main.go --cluster-size 5 --duration 30s --inject-failures
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type config struct {
	clusterSize     int
	duration        time.Duration
	numClients      int
	injectFailures  bool
	failureInterval time.Duration
	keySpace        int
	writeRatio      float64 // fraction of ops that are writes (vs. reads)
}

type result struct {
	ops       int64
	errors    int64
	latencies []time.Duration
	mu        sync.Mutex
}

func (r *result) record(d time.Duration, err error) {
	atomic.AddInt64(&r.ops, 1)
	if err != nil {
		atomic.AddInt64(&r.errors, 1)
		return
	}
	r.mu.Lock()
	r.latencies = append(r.latencies, d)
	r.mu.Unlock()
}

func main() {
	cfg := config{}
	flag.IntVar(&cfg.clusterSize, "cluster-size", 3, "number of raft nodes")
	flag.DurationVar(&cfg.duration, "duration", 30*time.Second, "how long to run the load test")
	flag.IntVar(&cfg.numClients, "clients", 10, "number of concurrent simulated clients")
	flag.BoolVar(&cfg.injectFailures, "inject-failures", false, "kill/restart the leader mid-run")
	flag.DurationVar(&cfg.failureInterval, "failure-interval", 10*time.Second, "how often to kill the leader when --inject-failures is set")
	flag.IntVar(&cfg.keySpace, "keyspace", 1000, "number of distinct keys to read/write")
	flag.Float64Var(&cfg.writeRatio, "write-ratio", 0.5, "fraction of operations that are writes")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "=== raft-kv-store benchmark ===\n")
	fmt.Fprintf(os.Stderr, "cluster size: %d | duration: %s | clients: %d | inject-failures: %v\n",
		cfg.clusterSize, cfg.duration, cfg.numClients, cfg.injectFailures)

	// TODO(you): replace this with real cluster bootstrapping once
	// kvraft1/ is implemented. This is where you'd:
	//   1. Spin up cfg.clusterSize raft1.Raft instances wired via labrpc
	//      (or real net/rpc if you're running an actual multi-process
	//      cluster instead of the in-process simulated one)
	//   2. Wrap them with your kvraft1 KV server
	//   3. Create cfg.numClients client handles pointed at the cluster
	log.Println("TODO: bootstrap cluster — see cmd/bench/main.go")

	res := &result{}
	var wg sync.WaitGroup
	stop := make(chan struct{})

	// Failure injector goroutine
	if cfg.injectFailures {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(cfg.failureInterval)
			defer ticker.Stop()
			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					// TODO(you): identify current leader and kill it,
					// then restart it after a random delay to simulate
					// a real crash/recovery cycle.
					log.Println("TODO: kill current leader")
				}
			}
		}()
	}

	// Simulated client load
	for i := 0; i < cfg.numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(int64(clientID)))
			for {
				select {
				case <-stop:
					return
				default:
				}
				start := time.Now()
				var err error
				// TODO(you): replace with real client.Get/Put calls once
				// the KV client is implemented.
				key := fmt.Sprintf("key-%d", rng.Intn(cfg.keySpace))
				if rng.Float64() < cfg.writeRatio {
					_ = key // client.Put(key, value)
				} else {
					_ = key // client.Get(key)
				}
				res.record(time.Since(start), err)
			}
		}(i)
	}

	time.Sleep(cfg.duration)
	close(stop)
	wg.Wait()

	printReport(res, cfg)
}

func printReport(r *result, cfg config) {
	r.mu.Lock()
	defer r.mu.Unlock()

	n := len(r.latencies)
	fmt.Printf("\n=== Results ===\n")
	fmt.Printf("total ops: %d | errors: %d\n", r.ops, r.errors)
	if n == 0 {
		fmt.Println("no successful ops recorded — implement client calls in cmd/bench/main.go")
		return
	}

	fmt.Printf("throughput: %.1f ops/sec\n", float64(n)/cfg.duration.Seconds())

	// naive percentile calc — fine for a benchmark script, swap for a real
	// histogram library if you want more precision
	sorted := append([]time.Duration(nil), r.latencies...)
	for i := 1; i < len(sorted); i++ {
		for j := i; j > 0 && sorted[j] < sorted[j-1]; j-- {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
		}
	}
	p50 := sorted[n*50/100]
	p99 := sorted[min(n-1, n*99/100)]
	fmt.Printf("p50 latency: %s | p99 latency: %s\n", p50, p99)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
