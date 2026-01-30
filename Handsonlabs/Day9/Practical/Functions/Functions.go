package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Funcitons Analysis Learning in Go")
	/*
		Creation and calling of Function Done
		Parameters multiple parameters with same type Done
		Return Types Done
		Mutliple return allowed Done
		Named (Naked ) return Done
		Error Done
		defer in go functions
		Single line parameter while sharing type Done
		Variadic Function Done
		Anonymous function Done
		Higher Order Function Done
		Call By Value and Call By Reference Done
	*/
	a := 500
	b := 200
	c := 400

	findGreaterNumber(a, b)
	name := "Alice"

	num, num1, name, sum := getNumbers(a, b, name)
	total := sumOfNumbers(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("Total Sum is", total)
	g := findGreaterNumber3(a, b, c)
	fmt.Println("Greater Number is", g)
	fmt.Println("Num1:", num1, " Num2:", num, "Name:", name, "Sum:", sum)
	studentMarks := 91
	grade := calCualteGrade(studentMarks)
	fmt.Println("Grade of Student is", grade)
	multificaitonOfTwoNumbers := higherOrderFunction(add, a, b)
	fmt.Println("Multiplication of Two Numbers is", multificaitonOfTwoNumbers)
	fmt.Println("Before Swapping A:", a, "B:", b)
	a, b = swapValues(&a, &b)
	fmt.Println("After Swapping A:", a, "B:", b)
	greatest, err := findGreaterNumber2(a, b)
	fileHandling("test.txt")
	if err != nil {
		fmt.Println("Error Occurred:", err)
	} else {
		fmt.Println("Greater Number is", greatest)
	}
}

//use of def in go functions

func fileHandling(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)

	} else {

		fmt.Println(file.Name())
	}
	defer file.Close() // Ensure the file is closed when the function exits
}

//fmt.Println("Control Statments")
// Find a greater number between 2 if else Done
// Find a greater number between 3  nest if else Done
// Find Grade of Student based on marks if else ladder in Done
// Switch Case for Print the Month  Name based on given number
// Try to write Grade program using switch
// For Loop
// print the table to given number
// Print values from 1 to 10
// Jumping Statement
// Continue and Break

/*fmt.Println("C",c)*/

func findGreaterNumber(a int, b int) {
	if a > b {
		fmt.Println("A", a)
	} else {
		fmt.Println("B", b)
	}
}

// Variadic Function
func sumOfNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Call By Reference using Pointers
func swapValues(a *int, b *int) (int, int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	return *a, *b
}

//error handling in go

func findGreaterNumber2(a int, b int) (int, error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
}

// func swapValues(a *int, b *int) {
// 	var temp int

// func findGreaterNumber3(a int, b int, c int) int {
// 	if a > b {
// 		if a > c {
// 			return a
// 		} else {
// 			return c
// 		}

// 	} else {
// 		if b > c {
// 			return b
// 		} else {
// 			return c
// 		}

// 	}
// }

func findGreaterNumber3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}

	} else {
		if b > c {
			return b
		} else {
			return c
		}

	}
}

//Named or Naked function Multiple return

/*func getNumbers(a,b int) (int,int){

	return a, b

}*/

//Named return or Naked function Multiple return

// func getNumbers() (num1 int, num2 int) {
// 	num1 = 10
// 	num2 = 20
// 	return
// }

func getNumbers(a, b int, name string) (num1 int, num2 int, name1 string, sum1 int) {
	num1 = 10 + a
	num2 = 20 + b
	name1 = name
	sum1 = num1 + num2
	return
}

// Higher Order Function
func higherOrderFunction(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func add(x, y int) int {
	return x + y
}

var calCualteGrade = func(studentMarks int) string {
	var grade string = ""
	if studentMarks > 90 {
		grade = "A"
	} else if studentMarks > 80 {
		grade = "B"
	} else if studentMarks > 70 {
		grade = "C"
	} else {
		grade = "D"
	}
	return grade
}

// if studentMarks > 90 {
// 	fmt.Println("Grade A")
// } else if studentMarks > 80 {

// 	fmt.Println("Grade B")

// } else if studentMarks > 70 {
// 	fmt.Println("Grade C")
// } else {
// 	fmt.Println("Grade D")
// }

// switch {
// case studentMarks > 90:
// 	fmt.Println("Grade A")
// case studentMarks > 80:
// 	fmt.Println("Grade B")
// case studentMarks > 70:
// 	fmt.Println("Grade C")
// default:
// 	fmt.Println("Grade D")

// }

// month := 2
// switch month {
// case 1:
// 	fmt.Println("Janunary")
// case 2:
// 	fmt.Println("Feb")
// case 3:
// 	fmt.Println("March")
// case 4:
// 	fmt.Println("April")
// case 5:
// 	fmt.Println("May")
// case 6:
// 	fmt.Println("Jun")
// case 7:
// 	fmt.Println("July")
// case 8:
// 	fmt.Println("August")
// case 9:
// 	fmt.Println("September")
// case 10:
// 	fmt.Println("October")
// case 11:
// 	fmt.Println("November")
// case 12:
// 	fmt.Println("December")
// default:
// 	fmt.Println("Wrong Entry")
// }

// //For Loop
// for i:=1; i<=10; i++ {
// 	fmt.Println(i)
// }

// //Print Table of Given Number Loop
// tableNumber :=2
// for i:=1; i<=10; i++ {
// 	fmt.Println(tableNumber*i)
// }

// for i:=1; i<=10; i++ {
// 	fmt.Println(tableNumber*i)
// 	if i==5 {
// 		 break
// 		}
// }

// for i:=1; i<=10; i++ {
// 	if i==5 {
// 	  continue
// 	}
// 	fmt.Println(tableNumber*i)
// }
