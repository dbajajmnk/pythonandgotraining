package main

import "fmt"

func main() {
	fmt.Println("Control Statments")
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
	a := 100
	b := 200
	c := 40
	/*fmt.Println("C",c)
	if a>b {
		fmt.Println("A",a)
	} else {
		fmt.Println("B",b)
	}
	*/

	if a > b {
		if a > c {
			fmt.Println("A is Greater", a)
		} else {
			fmt.Println("C is Greater", c)
		}

	} else {
		if b > c {
			fmt.Println("B is Greater", b)
		} else {
			fmt.Println("C is Greater", c)
		}

	}

	studentMarks := 91

	if studentMarks > 90 {
		fmt.Println("Grade A")
	} else if studentMarks > 80 {

		fmt.Println("Grade B")

	} else if studentMarks > 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	switch {
	case studentMarks > 90:
		fmt.Println("Grade A")
	case studentMarks > 80:
		fmt.Println("Grade B")
	case studentMarks > 70:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade D")

	}

	month := 2
	switch month {
	case 1:
		fmt.Println("Janunary")
	case 2:
		fmt.Println("Feb")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("Jun")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Wrong Entry")
	}

	//For Loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Print Table of Given Number Loop
	tableNumber := 2
	for i := 1; i <= 10; i++ {
		fmt.Println(tableNumber * i)
	}
	fmt.Println("Break Statement Example")
	for i := 1; i <= 10; i++ {
		fmt.Println(tableNumber * i)
		if i == 5 {
			break
		}
	}
	fmt.Println("Continue Statement Example")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(tableNumber * i)
	}

}
