package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email,omitempty"`
	Address Address `json:"address"`
}

func main() {
	user := User{
		ID:   1,
		Name: "Deepak",
		Address: Address{
			City:  "Delhi",
			State: "DL",
		},
	}

	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))

	f, _ := os.Create("user.json")
	defer f.Close()
	json.NewEncoder(f).Encode(user)
}
