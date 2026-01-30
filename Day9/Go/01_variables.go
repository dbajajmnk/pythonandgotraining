package main

import "fmt"

type Employee struct {
    ID   int
    Name string
}

func updateSalary(s *int) {
    *s = *s + 5000
}

func main(){

    // Primitive
    age := 30
    salary := 70000.50
    active := true
    name := "Deepak"

    // Composite
    scores := []int{85, 90}
    scores = append(scores, 95)

    subjects := map[string]int{"Go": 90, "Java": 85}

    emp := Employee{ID: 1, Name: "Deepak"}

    updateSalary(&age)

    fmt.Println(age, salary, active, name)
    fmt.Println(scores)
    fmt.Println(subjects)
    fmt.Println(emp)
}
