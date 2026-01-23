package main
import "fmt"

type User struct {
    name string
}

func createUser() *User {
    // Escape analysis decides heap allocation
    u := User{name: "Alice"}
    return &u
}

func main() {
    user := createUser()
    fmt.Println(user.name)
}
