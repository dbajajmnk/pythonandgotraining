package main

import (
	"fmt"
	"runtime"
)

/*
WHAT & WHY
Go compilation model is a major engineering advantage:
- go build produces a single executable binary
- fast builds (iteration speed)
- cross-compilation is easy (GOOS/GOARCH)

This program prints runtime + build hints.
*/

func main() {
	fmt.Println("=== GO COMPILATION MODEL ===")
	fmt.Printf("Runtime Go: %s\n", runtime.Version())
	fmt.Printf("Target OS/Arch at runtime: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	fmt.Println("\nBuild commands:")
	fmt.Println("1) Build current package:   go build")
	fmt.Println("2) Build with output name:  go build -o out/app")
	fmt.Println("3) Run without build file:  go run .")

	fmt.Println("\nCross-compile examples:")
	fmt.Println("GOOS=linux   GOARCH=amd64 go build -o out/app_linux_amd64 .")
	fmt.Println("GOOS=windows GOARCH=amd64 go build -o out/app_windows_amd64.exe .")
	fmt.Println("GOOS=darwin  GOARCH=arm64 go build -o out/app_macos_arm64 .")

	fmt.Println("\nEngineering note: single binary simplifies Docker + deployment.")
}
