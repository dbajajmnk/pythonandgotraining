package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
)

func main() {
    log.Println("pprof listening on :6060")
    log.Fatal(http.ListenAndServe(":6060", nil))
}
