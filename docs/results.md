# Benchmark Results

_(Fill in once Phase 7 is complete. Keep raw numbers + graphs here so you
can pull specifics into interviews without having to re-run everything.)_

## Setup

- Hardware: (your machine's specs)
- Cluster sizes tested: 3-node, 5-node
- Workload: (e.g. N concurrent clients, put/get ratio, key space size)

## Throughput / latency — healthy cluster

| Cluster size | Throughput (ops/sec) | p50 latency | p99 latency |
|---|---|---|---|
| 3 | | | |
| 5 | | | |

## Throughput / latency — under induced leader failure

| Cluster size | Downtime during re-election | Throughput after recovery | Notes |
|---|---|---|---|
| 3 | | | |
| 5 | | | |

## Takeaways

- (What was the actual trade-off you found? More nodes = more fault
  tolerance but how much replication overhead? Did leader failure recovery
  time match the ~5s the MIT tester expects?)
