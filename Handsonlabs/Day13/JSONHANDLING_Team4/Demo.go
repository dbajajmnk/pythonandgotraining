package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type car struct {
	COLOR   string `json:"color"`
	COMPANY string `json:"company"`
	ENGINE  engine `json:"engine"`
}

type engine struct {
	NUMBER string `json:"number"`
	POWER  string `json:"power"`
}

func main() {
	myCar := car{
		COLOR:   "Red",
		COMPANY: "BMW",
		ENGINE: engine{
			NUMBER: "10",
			POWER:  "100CC",
		},
	}

	data, _ := json.MarshalIndent(myCar, "", "  ")
	fmt.Println(string(data))
	file, err := os.Create("mycar.json")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(myCar)

}
