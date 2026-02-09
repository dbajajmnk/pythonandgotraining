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
