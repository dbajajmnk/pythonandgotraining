package main

import "fmt"

// func sum(a int,b int) int{
// 	return a+b;
// }
// func sum(a,b int) int{
// 	return a+b;
// }
//Multiple Return
func sum(a, b string) (string, string) {
	return a, b
}

//Naked and Named Return
func namedReturn(sum int) (a, b int) {
	a = sum * 10
	b = sum - 10
	return
}

//Variables Declaration
// var i,b,c,d bool
//Intitaliztion
// var a,b,c int = 10,20,40;
//With Different Values
// var a,b,c = "Deepak",true,1000;

//Short variable declarations applicable for function body only
//Zero values

func main() {
	var x = 50.5
	fmt.Println("Functions Examples")
	fmt.Println(sum("10", "20"))
	fmt.Println(namedReturn(18))
	fmt.Println(a, b, c, x)
}
