package main
import (
"fmt"
"sync"
"time"
)

func main(){
	fmt.Println("--------Goroutines in Go--------")
	go sayHelloWorld("Deepak")
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		sayHelloWorld("B")
	}()
	go func(){
		defer wg.Done()
		sayHelloWorld("C")
	}()

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("Loop:",i)
		}()
	}
	wg.Wait()
}
func sayHelloWorld(name string){
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello",name)

}
