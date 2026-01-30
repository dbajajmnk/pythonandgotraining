package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
WHAT & WHY
Control flow is how we "route" program execution:
- if / else
- switch
- for (Go has only one loop keyword)
- break/continue (including labeled)
- defer (run at function exit)

This file is a small interactive CLI to practice.
*/

func main() {
	fmt.Println("=== GO CONTROL FLOW (CLI LAB) ===")
	fmt.Println("Menu: 1) Grade 2) Sum 3) Search 0) Exit")

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nEnter choice: ")
		line, _ := in.ReadString('\n')
		choice := strings.TrimSpace(line)

		switch choice {
		case "1":
			gradeLab(in)
		case "2":
			sumLab(in)
		case "3":
			searchLab()
		case "0":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func gradeLab(in *bufio.Reader) {
	fmt.Print("Enter marks (0-100): ")
	line, _ := in.ReadString('\n')
	v := strings.TrimSpace(line)
	marks, err := strconv.Atoi(v)
	if err != nil || marks < 0 || marks > 100 {
		fmt.Println("Invalid marks")
		return
	}

	// if / else + switch demo
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	switch {
	case marks == 100:
		fmt.Println("Perfect score!")
	case marks%2 == 0:
		fmt.Println("Even score bonus note")
	default:
		fmt.Println("Odd score note")
	}
}

func sumLab(in *bufio.Reader) {
	defer fmt.Println("(defer) sumLab finished") // runs at function exit

	fmt.Print("Enter N to sum 1..N: ")
	line, _ := in.ReadString('\n')
	v := strings.TrimSpace(line)
	n, err := strconv.Atoi(v)
	if err != nil || n < 1 {
		fmt.Println("Invalid N")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ { // classic for
		sum += i
	}
	fmt.Println("Sum:", sum)
}

func searchLab() {
	// labeled break/continue demo
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	target := 5

	found := false

Outer:
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == target {
				fmt.Printf("Found %d at [%d][%d]\n", target, i, j)
				found = true
				break Outer
			}
		}
	}

	if !found {
		fmt.Println("Not found")
	}
}
