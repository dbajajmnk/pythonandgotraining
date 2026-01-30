/******* Basic Structure *********************** Haripriya/
/********* Primitive Types********************* Manoj/
/*********** Zero Values ******************** Manoj/
/*********** Verobcity ********************* Harivinder/
/************ Single line Declaration and Try multiple values**** Harivinder/
/*********** Composite Types **************/
/************ Array of Different Types*****/
/****** Struct ********************/
/** Map *********************/
/** Slice **************/
/*** Pointer implemenation ******/
package main

import "fmt"

//outside : = "Harivardan"
var outsider="Deepak"
func main() {

	fmt.Println("Hello World")
	fmt.Println(outsider)
	//fmt.Println(outside)
/*
	// Prmitive Types
//Zero values
	var age int
	var isAdult bool
	var spends float64
	var name string

	fmt.Println("Name is ", name)
	fmt.Println("Is Adult ", isAdult)
	fmt.Println("Age is ", age)
	fmt.Println("Expense", spends)
//Initialize Values
	age = 23
	isAdult = true
	spends = 2000.54
	name = "Manoj"

	fmt.Println("Name is ", name)
	fmt.Println("Is Adult ", isAdult)
	fmt.Println("Age is ", age)
	fmt.Println("Expense", spends)
*/
// Prmitive Types
//Verbocity

	age := 23
	isAdult := true
	spends := 2000.54
	name := "Manoj"

	fmt.Println("Name is ", name)
	fmt.Println("Is Adult ", isAdult)
	fmt.Println("Age is ", age)
	fmt.Println("Expense", spends)

//Single line Verbocity 
a,b,c := 10,20,30
f,g,h := true,false,true
k,l,m :="K","L", "M"
p,q,r := 10.0,4.5,2.909

fmt.Println(a,b,c,f,g,h,k,l,m,p,q,r)

var title,price,page,author,isAvailable = "Go Lang",100.50,500,"Harivardan",true

fmt.Println(title,price,page,author,isAvailable)

//Comosite Types 
var students[3]string = [3]string{"Deepak","Sunil","Manoj"}

fmt.Println(students)

//Slice
first2 := students[:2]
fmt.Println("Slice",first2)

type Student struct{
   ID int 
   rollNo int
   name string 
   stream string

}

var Harivardan Student = Student{ID:1,rollNo:10, name:"Harivardan", stream :"Btech"}
fmt.Println(Harivardan);
//Map 
classResult :=map[string]float32 {"Ram":90.6,"Sham":89,"John":20}
fmt.Println(classResult)


}
