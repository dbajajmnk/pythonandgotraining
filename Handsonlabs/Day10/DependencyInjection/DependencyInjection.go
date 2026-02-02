package main

import "fmt"

/**
Health  App
App
User Admin
Staff
Patient
// Composition over Inheritance
// Use many Struct are required here
// Which things are going to depency injection
// Which attaching method in struct









*/

type Feeback interface {
	submitFeedback()
}

type printUserInfo interface {
	printInfo()
}

type User struct {
	Name    string
	Age     int
	Address string
	Phone   string
	Email   string
}

func (u User) printInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, Phone: %s, Email: %s\n", u.Name, u.Age, u.Address, u.Phone, u.Email)
}

type Admin struct {
	User
	AdminLevel int
}
type Staff struct {
	User
	Department string
}

func (s Staff) submitFeedback() {
	fmt.Printf("Staff %s is submitting feedback for department %s\n", s.Name, s.Department)
}

func (p Patient) submitFeedback() {
	fmt.Printf("Patient %s is submitting feedback for medical record: %s\n", p.Name, p.MedicalRecord)
}

type Patient struct {
	User
	MedicalRecord string
}

func main() {
	fmt.Println("Dependency Injection in Go")
	admin := Admin{
		User: User{
			Name:    "Alice",
			Age:     35,
			Address: "123 Admin St",
		},
		AdminLevel: 1,
	}

	staff := Staff{
		User: User{
			Name:    "Bob",
			Age:     30,
			Address: "456 Staff Ave",
		},
		Department: "Cardiology",
	}

	patient := Patient{
		User: User{
			Name:    "Charlie",
			Age:     25,
			Address: "789 Patient Rd",
		},
		MedicalRecord: "Healthy",
	}

	admin.printInfo()
	staff.printInfo()
	patient.printInfo()
	patient.submitFeedback()
	staff.submitFeedback()

}
