.PHONY: test-3a test-3b test-3c test-3d test-all bench bench-failures fmt vet

test-3a:
	cd raft1 && go test -v -race -run 3A

test-3b:
	cd raft1 && go test -v -race -run 3B

test-3c:
	cd raft1 && go test -v -race -run 3C

test-3d:
	cd raft1 && go test -v -race -run 3D

test-all: test-3a test-3b test-3c test-3d

bench:
	go run cmd/bench/main.go --cluster-size 3 --duration 30s
	go run cmd/bench/main.go --cluster-size 5 --duration 30s

bench-failures:
	go run cmd/bench/main.go --cluster-size 3 --duration 30s --inject-failures
	go run cmd/bench/main.go --cluster-size 5 --duration 30s --inject-failures

fmt:
	gofmt -l -w .

vet:
	go vet ./...
