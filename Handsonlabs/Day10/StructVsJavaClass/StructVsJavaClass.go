package main
import (
	"fmt"
)
type Student struct {
	Name string
	Age  int
}

func main() {
	student1 := Student{Name: "Alice", Age: 21}
	student2 := Student{Name: "Bob", Age: 22}		
	fmt.Println("Student 1:", student1)
	fmt.Println("Student 2:", student2)
}