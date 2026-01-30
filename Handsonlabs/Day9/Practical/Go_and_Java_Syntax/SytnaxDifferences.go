/*
Hello World
Basic variable with Basic values
if condition
variable declaration and initialization of same type in single line
showing ; is optional in Go
Comments in Go
Naming Conventions in Go what we understand
:a=10 varbosity
*/

package main 
import "fmt"

func main(){
	

fmt.Println("Hello World") // Hello World 
//var a int 
a := 20 // Basic variable with Basic values
x, y, z := 1, 2, 3 // variable declaration and initialization of same type in single line

sum := x + y + z
fmt.Println("Sum is:", sum)
//Showing ; is optional in Go
var isAdult bool 

//Above variable with basic values
var one,two,three,four,five = 1,2,3,4,5
//variable declaration and initialization in same line
//var age int = 20

fmt.Println(one,two,three,four,five)

fmt.Println(a)
fmt.Println(isAdult)
//Example of if condition
if age>18 {
	fmt.Println("You are Eligible for Voting")

}





}