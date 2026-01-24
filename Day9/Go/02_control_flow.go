package main
import "fmt"

func controlFlow() {
    // if / else
    x := 5
    if x > 3 {
        fmt.Println("x > 3")
    }

    // for loop
    for i := 0; i < 3; i++ {
        fmt.Println("i =", i)
    }
}
