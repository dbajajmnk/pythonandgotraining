package main

import "fmt"

func main() {
    // Array (fixed size)
    var a [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", a)

    // Slice (dynamic)
    s := []int{1, 2, 3}
    s = append(s, 4)
    fmt.Println("Slice:", s, "len:", len(s), "cap:", cap(s))

    // Map (key-value)
    m := map[string]int{
        "apple":  2,
        "banana": 5,
    }

    // Safe access
    val, ok := m["apple"]
    fmt.Println("apple:", val, "exists:", ok)

    _, exists := m["orange"]
    fmt.Println("orange exists:", exists)
}
