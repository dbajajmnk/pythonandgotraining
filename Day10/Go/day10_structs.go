package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Deepak", Age: 30}
	fmt.Println(u)
}
