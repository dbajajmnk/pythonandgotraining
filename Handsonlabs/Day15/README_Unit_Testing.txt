Unit Testing Lab (Go)

Folder structure:
  unit_testing_demo/
    go.mod
    calculator.go
    calculator_test.go

Steps:
  1) mkdir unit_testing_demo && cd unit_testing_demo
  2) go mod init unit_testing_demo
  3) Create files exactly as provided.
  4) Run:
       go test -v
       go test -cover

Notes:
  - Tests must be in *_test.go
  - Package name can be 'calculator' as shown.


Default Value of b.N
There is no fixed default value for b.N. It starts very small (usually 1) and grows automatically in a loop (1, 2, 5, 10, 20, 50, 100...) until the total runtime meets the -benchtime requirement. 
How to Change b.N
You can change the number of iterations or the duration of the benchmark using go test flags: 
Change b.N (Total Iterations):
Use -benchtime with an "x" suffix to force a specific number of iterations.
bash
# Run the benchmark exactly 100 times
go test -bench=. -benchtime=100x
Change Time Duration (Default 1s):
Change the total time the benchmark runs to allow b.N to grow larger.
bash
# Run the benchmark for 10 seconds instead of 1 second
go test -bench=. -benchtime=10s
