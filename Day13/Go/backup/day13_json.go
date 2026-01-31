package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    u := User{Name: "Deepak", Age: 30}

    data, _ := json.Marshal(u)
    fmt.Println(string(data))

    var u2 User
    json.Unmarshal(data, &u2)
    fmt.Println(u2)
}
