/**
Single Return
Multiple Return
	1. Result+Error
	2. Result+bool (Lookups)
	3. Grouping (min,max) context
	4. Result + metadata
Named Resturns
Defer wrapping
Callback
Factory pattern (return faunction from another function)
**/

package main

import (
	"fmt"
)

func main() {
	result,err := divideByZero(20,0)
	fmt.Println("Divide By Zero",result,err)
	
	
	
}
//Single Return
func add(x, y int) int {
	return x + y
}
//Multiple Return Return+Error
func divideByZero(value,divideNumber int) (float64,error){
  if divideNumber==0 {
	return 0,fmt.Errorf("You can't divide the number by %.02f",float64(divideNumber))
	}
   return float64(value)/float64(divideNumber),nil 
    
}
//Lookup (Exist) Return+ok
func isAvailable(students map[string]string,uuid string)(string,bool){
	value,ok := students[uuid]
	return value,ok
}

//Group (Based on context)
func getGreaterAndSmaller(num1,num2 int) (int,int) {
	min,max :=0,0
	if(num1>num2){
		max =num1
		min =num2

	}else {

		max =num2
		min =num1
	}
  return max,min
} 

func metaDataWithReturn(students map[string]string) (string,int) {
	
	return students["1"],len(students)

}

