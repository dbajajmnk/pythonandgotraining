package main

import "fmt"

type Car interface {
    drive() string
	applyBreak()
}

type Engine interface {
    start() string
	stop() 
}

type Tesla struct { Model string }
func (t Tesla) start() string { return "Starting Tesla " + t.Model }
func (t Tesla) stop() { fmt.Println("Stopping Tesla") }
func (t Tesla) drive() string { return "Driving a Tesla " + t.Model }
func (t Tesla) applyBreak() { fmt.Println("Applying break for Tesla") }
type BMW struct { Series string }
func (b BMW) drive() string { return "Driving a BMW " + b.Series }
func (b BMW) applyBreak() { fmt.Println("Applying break for BMW") }

type Nexon struct { Variant string }
func (n Nexon) drive() string { return "Driving a Nexon " + n.Variant }
func (n Nexon) applyBreak() { fmt.Println("Applying break for Nexon") }


func main() {
    cars := []Car{Tesla{Model: "Model S"}, BMW{Series: "X5"}, Nexon{Variant: "XUV700"}}
    for _, s := range cars {
        fmt.Println(s.drive())
        s.applyBreak()
    }
}