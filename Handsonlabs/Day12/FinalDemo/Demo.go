package main

import (
	"fmt"
)

func main() {
	printWrapper("Channels lession Start")
	numCh := make(chan int)
	go func() {
		for i := range 5 {
			numCh <- i

		}
		close(numCh)

	}()

	for data := range numCh {

		printWrapper("Value", data)
	}

	directional := make(chan bool)
	go producer(directional)
	<-directional

	printWrapper("Channel Lab is Done")

	// // bmw := BMW{
	// // 	COMPANY: "BMW",
	// // 	WHEELSCOUNT: 4,
	// // }
	// nexon := NEXON{
	// 	A:"NEXON",
	// 	B:4,

	// }
	// showCarFeature(nexon)

}
func producer(done chan<- bool) {
	printWrapper("Producer is active")
	done <- true

}
func senderForString(done chan<- string, data string) {
	done <- data
}
func printWrapper(data ...any) {
	// newValue := append(data,"posfix")
	
	// fmt.Println(newValue...)

	//  fmt.Println(append(append([]any{"WOW"}, data...), "posfix")...)
	fmt.Println(append([]any{"WOW"}, data...)...)
}

// type BMW struct {
// 	COMPANY string
// 	WHEELSCOUNT int
// }

// type NEXON struct {
// 	A string
// 	B int
// }
// type PUNCH struct {
// 	C string
// 	D int
// }
// type SCORPIO struct {
// 	E string
// 	F int
// }

// type carFeatures interface {
// 	showFeatures()
// }

// func (bm BMW) showFeatures(){
// 	fmt.Println("BMW",bm.COMPANY)
// 	fmt.Println("BMW",bm.WHEELSCOUNT)
// }

// func (bm NEXON) showFeatures(){
// 	fmt.Print(bm.A)
// 	fmt.Println(bm.B)
// }
// func (bm PUNCH) showFeatures(){
// 	fmt.Print(bm.C)
// 	fmt.Println(bm.D)
// }
// func (bm SCORPIO) showFeatures(){
// 	fmt.Print(bm.E)
// 	fmt.Println(bm.F)
// }

// func showCarFeature(car carFeatures){
// 	car.showFeatures()
// }


// func doWork(){
// 	time.Sleep(200*time.Millisecond)
// 	fmt.Print("I am Called")
// }

