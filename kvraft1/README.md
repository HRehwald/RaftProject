Empty — copy from mit-6.5840/src/kvraft1/ once you reach that phase.
This is the key-value service built on top of raft1/. It's a later
lab (Lab 4), optional for the scope of this project — the KV layer
can also be hand-rolled directly on raft1's Start()/ApplyMsg interface
if you want to keep the project scope to just Raft + a thin KV wrapper.
