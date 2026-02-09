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
		doWork()
	}
}
func doWork() {
	var sum int
	for i := 0; i < 100_000_000; i++ {
		sum += i
	}
	time.Sleep(100 * time.Microsecond)
	_ = sum

}
