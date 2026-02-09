// File: mathops.go
package mathops

func Add(a, b int) int {
    return a + b
}

func AddWithLoop(a, b int) int {
    sum := 0
    for i := 0; i < b; i++ {
        sum += 1
    }
    return a + sum - b
}