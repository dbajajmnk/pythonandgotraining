package main

import "fmt"

func main() {
    // make for slice
    s := make([]int, 0, 5)
    s = append(s, 1, 2, 3)
    fmt.Println("Slice:", s, "len:", len(s), "cap:", cap(s))

    // copy
    dst := make([]int, len(s))
    copy(dst, s)
    fmt.Println("Copied slice:", dst)

    // delete map entry
    m := map[string]int{"a": 1, "b": 2}
    delete(m, "a")
    fmt.Println("Map after delete:", m)
}
