package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
WHAT & WHY
Go is built for "engineering throughput":
- Fast compile times (developer velocity)
- Simple language surface (less cognitive load)
- Great concurrency model (goroutines + channels)
- Strong standard library (net/http, json, testing)

This file demonstrates:
1) A tiny "service-like" workload (CPU + I/O simulation)
2) Concurrency using goroutines
3) Predictable output & easy deploy (single binary)
*/

func main() {
	fmt.Println("=== WHY GO (Engineering Demo) ===")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	// Mind map in code: "We have tasks coming in; handle them concurrently"
	tasks := []Task{
		{ID: "T1", WorkMs: 250},
		{ID: "T2", WorkMs: 120},
		{ID: "T3", WorkMs: 400},
		{ID: "T4", WorkMs: 180},
	}

	start := time.Now()
	results := processTasksConcurrently(tasks, 2) // 2 workers
	elapsed := time.Since(start)

	fmt.Println("\n--- Results ---")
	for _, r := range results {
		fmt.Printf("%s => %s (took %v)\n", r.TaskID, r.Status, r.Duration)
	}
	fmt.Printf("\nTotal elapsed: %v\n", elapsed)
	fmt.Println("\nNext step: `go build` this program -> you get a single binary to ship.")
}

type Task struct {
	ID     string
	WorkMs int
}

type Result struct {
	TaskID    string
	Status    string
	Duration  time.Duration
	Completed bool
}

// Engineering: Worker pool pattern (common in services)
func processTasksConcurrently(tasks []Task, workers int) []Result {
	taskCh := make(chan Task)
	resultCh := make(chan Result)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for t := range taskCh {
				r := doWork(workerID, t)
				resultCh <- r
			}
		}(i + 1)
	}

	// Feed tasks
	go func() {
		for _, t := range tasks {
			taskCh <- t
		}
		close(taskCh)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect
	var results []Result
	for r := range resultCh {
		results = append(results, r)
	}
	return results
}

// Simulated work (CPU + waiting)
func doWork(workerID int, t Task) Result {
	start := time.Now()
	// "Work" = sleep to simulate I/O. In real systems: DB call, HTTP call, etc.
	time.Sleep(time.Duration(t.WorkMs) * time.Millisecond)

	return Result{
		TaskID:    t.ID,
		Status:    fmt.Sprintf("DONE by worker-%d", workerID),
		Duration:  time.Since(start),
		Completed: true,
	}
}
