package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("==========Go Routines===========")
	go sayHello("A")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sayHello("B")
	}()
	go func() {
		defer wg.Done()
		sayHello("C")
	}()

	studentNames := [3]string{"Ram", "John", "Tom"}

	// for i:=0; i<len(studentNames);i++ {
	// 	wg.Add(1)
	// 	go func(){
	// 		defer wg.Done()
	// 		fmt.Println("Bravo !", studentNames[i])
	// 	}()

	// }

	for index, name := range studentNames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Well Done",index, name)
		}()
	}

	wg.Wait()
}
func sayHello(name string) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Hello ", name)

}
