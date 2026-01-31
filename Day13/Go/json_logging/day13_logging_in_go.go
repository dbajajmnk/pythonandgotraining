// File: day13_logging_in_go.go
// Topic: Logging in Go (Day 13)
// Run: go run day13_logging_in_go.go
//
// Demonstrates:
// 1) standard log package
// 2) logging to file
// 3) structured logging using slog (JSON)
// 4) correlation/request_id pattern
// 5) safe logging (mask secrets)

package main

import (
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	// ----- Lab 1 & 2: Basic log + flags -----
	log.SetPrefix("[APP] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("starting application")

	// ----- Lab 3: Log to file -----
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("failed to open log file:", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("this line goes to app.log")

	// ----- Lab 4: slog JSON logs -----
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)

	// ----- Lab 5: Correlation ID -----
	requestID := "req-12345"
	reqLogger := logger.With("request_id", requestID)

	reqLogger.Info("incoming request", "path", "/users", "method", "GET")

	// ----- Lab 6: Error with context -----
	if err := callExternalService(); err != nil {
		reqLogger.Error("external call failed", "err", err, "retry", false)
	}

	// ----- Lab 7: Safe logging (mask secrets) -----
	token := "secret-token-abc"
	reqLogger.Warn("token received (masked)", "token", mask(token))

	reqLogger.Info("done")
}

func callExternalService() error {
	time.Sleep(100 * time.Millisecond)
	return os.ErrNotExist
}

func mask(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:2] + "****" + s[len(s)-2:]
}
