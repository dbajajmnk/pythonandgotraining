Unit Testing Lab (Go)

Folder structure:
  benchmark/
  mathops
    go.mod
    .go
    mathops.go

Steps:
  1) mkdir benchmark && cd benchmark
  2) go mod init benchmark
  3) Create files exactly as provided.
  4) Run:
        go test -bench=.
        go test -bench=. -count 230
        go test -bench=. -count 230 -run=^#

Notes:
  - Tests must be in *_test.go
  - Test fuction name prefix with Test
  - Bench mark funcitons prefix with Benchmark
  - Package name can be 'mathops' as shown.
