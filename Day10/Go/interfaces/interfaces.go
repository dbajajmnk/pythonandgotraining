package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

type Rectangle struct{ W, H float64 }
func (r Rectangle) Area() float64 { return r.W * r.H }

func main() {
    shapes := []Shape{Circle{2}, Rectangle{3,4}}
    for _, s := range shapes {
        fmt.Println(s.Area())
    }
}
