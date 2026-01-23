package main

import "fmt"

type Printer interface {
	Print()
}

type Report struct{}

func (r Report) Print() {
	fmt.Println("Printing report")
}

func main() {
	var p Printer = Report{}
	p.Print()
}
