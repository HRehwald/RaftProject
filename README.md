# raft-kv-store

A fault-tolerant, replicated key-value store built on the Raft consensus
algorithm, implemented in Go. Built on top of the MIT 6.5840 (Distributed
Systems) lab framework, extended with a fault-injection test suite,
throughput/latency benchmarking, and (stretch) linearizability checking.

## Why this project

Raft is a well-understood consensus algorithm, which means the interesting
engineering work isn't "does it work at all" — it's proving *how* it behaves
under real failure conditions (crashed nodes, network partitions, message
loss/reordering) and *how fast* it is under different cluster shapes. That's
what the `bench/` and `docs/results.md` pieces of this repo are for, on top
of the core Raft + KV implementation.

## Repo layout

```
raft-kv-store/
├── raft1/          Core Raft implementation (leader election, log
│                   replication, persistence, snapshotting)
├── raftapi/        The interface Raft exposes to the service layer above it
├── kvraft1/         Key-value service built on top of raft1/
├── labrpc/         Simulated RPC layer used by the tester (delays, drops,
│                   reorders messages to simulate real network conditions)
├── labgob/         Wrapper around Go's gob encoder with stricter checks
├── tester1/        Test harness: cluster setup, fault injection, the
│                   official correctness test suite
├── main/           Entry points for running a real (non-simulated) cluster
├── cmd/bench/      Benchmarking harness: throughput/latency under load
│                   and under induced failures
└── docs/
    ├── PROJECT_PLAN.md   The 8-phase build plan (scope -> election ->
    │                     replication -> persistence -> snapshotting ->
    │                     fault injection -> benchmarking -> stretch goal)
    └── results.md        Where benchmark results + graphs go once you have
                          real data (this is the file that becomes your
                          resume bullet and interview talking point)
```

## Getting the official MIT lab code

`raft1/`, `raftapi/`, `labrpc/`, `labgob/`, and `tester1/` are currently
empty placeholders. Pull the real starter code + official test suite from
MIT's 6.5840 course (this must be run from your own machine — MIT
distributes it over `git://`, which most sandboxed/cloud dev environments
block):

```bash
git clone git://g.csail.mit.edu/6.5840-golabs-2026 mit-6.5840
```

If that specific tag has rotated by the time you read this, check the
current year's tag on the course's "Getting Started" page:
https://pdos.csail.mit.edu/6.824/labs/lab-mr.html (the git clone command is
under "Getting Started" — the year in the URL/tag changes each term).

Then copy the pieces you need into this repo, overwriting the empty
placeholders:

```bash
cp -r mit-6.5840/src/raft1/*      raft-kv-store/raft1/
cp -r mit-6.5840/src/raftapi/*    raft-kv-store/raftapi/
cp -r mit-6.5840/src/labrpc/*     raft-kv-store/labrpc/
cp -r mit-6.5840/src/labgob/*     raft-kv-store/labgob/
cp -r mit-6.5840/src/tester1/*    raft-kv-store/tester1/
cp -r mit-6.5840/src/kvraft1/*    raft-kv-store/kvraft1/   # once you reach that phase
```

You'll now have the real `raft.go` skeleton (with `Make()`, `Start()`,
`GetState()` stubs and RPC struct definitions) plus the official test suite
in `raft_test.go`. **Read the extended Raft paper's Figure 2 before writing
any code** — https://pdos.csail.mit.edu/6.824/papers/raft-extended.pdf.

## Running tests

Once the real code is copied in:

```bash
cd raft1
go test -v -race -run 3A   # leader election
go test -v -race -run 3B   # log replication
go test -v -race -run 3C   # persistence
go test -v -race -run 3D   # snapshotting / log compaction
```

## Running the benchmark suite (your own work, not MIT's)

This is the part that's actually yours — see `docs/PROJECT_PLAN.md` phase 7
and `cmd/bench/`.

```bash
go run cmd/bench/main.go --cluster-size 3 --duration 30s
go run cmd/bench/main.go --cluster-size 5 --duration 30s --inject-failures
```

Results and write-up go in `docs/results.md`.

## License / academic integrity note

The `raft1/`, `raftapi/`, `labrpc/`, `labgob/`, and `tester1/` packages,
once populated, are MIT's course materials — used here for personal,
non-credit learning, not for submission to any course. Don't publish
completed solutions if you ever *are* enrolled in 6.5840 or an equivalent
course elsewhere; check your course's collaboration policy first.
