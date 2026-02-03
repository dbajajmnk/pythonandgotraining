// /*
// *
// Validation Error Handling
// API Error Handling

// 	No Data Found
// 	Server Error
// 	Invalidate Input
// 	Missing Parameter

// Passing Error to another layer
// Panic
// *
// */
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	fmt.Println("=== Exception Handling in Go ===")
// 	p := Person{
// 		Name:  "Deepak",
// 		Age:   19,
// 		Phone: "89283838838",
// 		Email: "test@example.com"}

// 	// err := validatePerson(p)
// 	// if err != nil {
// 	// 	fmt.Println("Validation Error:", err)
// 	// } else {
// 	// 	fmt.Println("Person is valid:", p)
// 	// }
// 	// response, err := fetchData(p, "My Student are Great")
// 	// if err != nil {
// 	// 	fmt.Println("API Error:", err)
// 	// } else {
// 	// 	fmt.Println("Fetched Data:", response)
// 	// }
// 	panicExample(p, "No Data")
// 	fmt.Println("=== Done ===")

// }

// type Person struct {
// 	Name  string
// 	Age   int
// 	Phone string
// 	Email string
// }

// // Validation Error Handling
// func validatePerson(p Person) error {
// 	if p.Name == "" {
// 		return errors.New("name cannot be empty")
// 	} else if p.Age < 0 {
// 		return errors.New("age cannot be negative")
// 	} else if p.Age < 18 {
// 		return errors.New("age must be at least 18")
// 	} else if p.Phone == "" {
// 		return errors.New("phone cannot be empty")
// 	} else if p.Email == "" {
// 		return errors.New("email cannot be empty")
// 	}
// 	return nil
// }

// func fetchData(request Person, dummyResponse string) (Person, error) {
// 	err := validatePerson(request)
// 	if err != nil {
// 		return Person{}, fmt.Errorf("fetchData failed: %w", err)
// 	}

// 	switch dummyResponse {
// 	case "Server Error":
// 		return Person{}, errors.New("internal server error")
// 	case "No Data":
// 		return Person{}, errors.New("no data found")

// 	}

// 	// Simulate data fetching
// 	return request, nil

// }

// func panicExample(request Person, dummyResponse string) {
// 	defer func() {
// 		if response, r := fetchData(request, dummyResponse); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 		} else {
// 			fmt.Println("Fetched Data:", response)
// 		}
// 		fmt.Println("This will not be printed")

// 	}()
// 	panic("Unexpected error occurred")
// }


package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("=== Exception Handling in Go ===")
	p := Person{
		Name:  "Deepak",
		Age:   19,
		Phone: "89283838838",
		Email: "test@example.com"}

	// err := validatePerson(p)
	// if err != nil {
	// 	fmt.Println("Validation Error:", err)
	// } else {
	// 	fmt.Println("Person is valid:", p)
	// }
	// response, err := fetchData(p, "My Student are Great")
	// if err != nil {
	// 	fmt.Println("API Error:", err)
	// } else {
	// 	fmt.Println("Fetched Data:", response)
	// }
	panicExample(p, "No Data")
	fmt.Println("=== Done ===")

}

type Person struct {
	Name  string
	Age   int
	Phone string
	Email string
}

// Validation Error Handling
func validatePerson(p Person) error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	} else if p.Age < 0 {
		return errors.New("age cannot be negative")
	} else if p.Age < 18 {
		return errors.New("age must be at least 18")
	} else if p.Phone == "" {
		return errors.New("phone cannot be empty")
	} else if p.Email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}

func fetchData(request Person, dummyResponse string) (Person, error) {
	err := validatePerson(request)
	if err != nil {
		return Person{}, fmt.Errorf("fetchData failed: %w", err)
	}

	switch dummyResponse {
	case "Server Error":
		return Person{}, errors.New("internal server error")
	case "No Data":
		return Person{}, errors.New("no data found")

	}

	// Simulate data fetching
	return request, nil

}

func panicExample(request Person, dummyResponse string) {
	defer func() {
		if response, r := fetchData(request, dummyResponse); r != nil {
			fmt.Println("Recovered from panic:", r)
		} else {
			fmt.Println("Fetched Data:", response)
		}
		fmt.Println("This will not be printed")

	}()
	panic("Unexpected error occurred")
}
