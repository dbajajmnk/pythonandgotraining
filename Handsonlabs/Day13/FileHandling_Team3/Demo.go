package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Handling")
	file, error := os.Create("team3.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString("Team 3 Fist Task Done\n")

	file2, err := os.OpenFile("team3.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(error)
	}
	defer file2.Close()
	file2.WriteString("I am writing in Second File\n")

	buffer := bufio.NewWriter(file2)
	buffer.WriteString("I am writting with Buffer\n")
	buffer.Flush()

}
