// File: day15_profiling_basics.go
// Topic: Profiling Basics in Go
// Run:
//   go run day15_profiling_basics.go
//   go tool pprof http://localhost:6060/debug/pprof/profile
//   go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println("pprof server running on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		busyWork()
	}
}

func busyWork() {
	sum := 0
	for i := 0; i < 1_000_000; i++ {
		sum += i
	}
	time.Sleep(100 * time.Millisecond)
	_ = sum
}
