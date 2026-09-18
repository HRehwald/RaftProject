# Project Plan

Tracks the 8-phase build plan. Check boxes off as you go — this doubles as
your progress log and, later, the source material for resume bullets and
interview talking points.

## Phase 1 — Scope it before writing code
- [ ] Write a one-paragraph spec: supported ops (get/put/delete), target
      consistency guarantee (linearizable reads/writes), explicit
      non-goals (multi-datacenter, sharding, auth)
- [ ] Paste the spec into this repo's top-level README intro once decided

## Phase 2 — Leader election (Raft paper §5.1–5.2, Figure 2)
- [ ] Randomized election timeouts
- [ ] RequestVote RPC (sender + handler)
- [ ] Term handling / conversion rules (follower -> candidate -> leader)
- [ ] Passes MIT's 3A test suite (`go test -run 3A`)
- [ ] Manually verify: kill the leader, confirm re-election within ~5s

## Phase 3 — Log replication (Raft paper §5.3, Figure 2)
- [ ] AppendEntries RPC (sender + handler)
- [ ] Log matching property maintained under out-of-order/duplicate RPCs
- [ ] Commit index advancement
- [ ] Passes MIT's 3B test suite (`go test -run 3B`)

## Phase 4 — Persistence + client API (Raft paper §5.5, §8)
- [ ] Persist currentTerm, votedFor, log to the Persister
- [ ] Restore state correctly on restart
- [ ] get/put/delete client API routed through current leader
- [ ] Client retries on leader-redirect / timeout
- [ ] Passes MIT's 3C test suite (`go test -run 3C`)

## Phase 5 — Snapshotting / log compaction (Raft paper §7)
- [ ] Periodic snapshots, discarding log entries before snapshot index
- [ ] InstallSnapshot RPC for lagging followers
- [ ] Passes MIT's 3D test suite (`go test -run 3D`)

## Phase 6 — Fault injection test harness (your own work)
- [ ] Test: kill a node mid-write, confirm cluster still converges
- [ ] Test: partition the network (drop messages between subsets), confirm
      no split-brain and correct recovery on heal
- [ ] Test: restart nodes with stale/missing state, confirm correct catch-up
- [ ] Write a short doc (`docs/fault-injection-notes.md`) describing each
      scenario and what you verified — this is the "prove it" deliverable

## Phase 7 — Benchmark + write-up (your own work)
- [ ] Measure throughput/latency at 3-node vs 5-node cluster size
- [ ] Measure throughput/latency under induced leader failure mid-load
- [ ] Graph results (`cmd/bench/` -> `docs/results.md`)
- [ ] Write a short results doc: what you measured, what you expected,
      what you found, and why (this becomes your resume bullet)

## Phase 8 — Stretch: linearizability checking
- [ ] Record operation histories from your fault-injection tests
- [ ] Run histories through Porcupine (MIT's tests already use this —
      check `tester1/` once pulled in) or a standalone checker
- [ ] Note any violations found and how you fixed them

---

## Resume bullet draft (fill in once phases 6–7 are done)

> Built a fault-tolerant, replicated key-value store on the Raft consensus
> algorithm in Go, implementing leader election, log replication,
> persistence, and log compaction from the extended Raft paper. Validated
> correctness under induced node crashes and network partitions with a
> custom fault-injection test suite, and benchmarked throughput/latency
> across [N]-node clusters under both healthy and degraded conditions.
