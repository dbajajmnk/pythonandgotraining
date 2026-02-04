package main
import (
"fmt"
"time"
)

func main(){
	fmt.Println("--------Goroutines in Go--------")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		ch1<- "Message 1"
	}()
	go func(){
		ch2<- "Message 2"
	}()
	
	for i:=0;i<2;i++ {
		select {
		case msg:= <-ch1:
			fmt.Println("Message One Recieved",msg,value)
		case msg:= <-ch2:
			fmt.Println("Message One Recieved",msg,value)
		}

	}
}
func sayHelloWorld(name string){
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello",name)

}
