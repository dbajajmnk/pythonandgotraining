package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Collection and Inbuild Methods ===")
	//Array
	array := [3]string{"One","Two","Three"}
	
	
	//Slice
	fmt.Println("Array",array, len(array),cap(array))
	slice := []string{"One","Two","Three"}
	fmt.Println("Slice",slice)

	
	m := map[string]int {"a":1,"b":2}
	fmt.Println("Map",m)
	map2 := m
	
	map2["c"]= 3
	fmt.Println("Map 2",map2)
	fmt.Println("Map",m)






	

}

