package main

import "fmt"

/**
Health  App
App
User 
Admin
Staff
Patient
Driver
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

func (u Admin) printInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, Phone: %s, Email: %s, AdminLevel: %d\n", u.Name, u.Age, u.Address, u.Phone, u.Email, u.AdminLevel)
}

func (u Staff) printInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, Phone: %s, Email: %s, Department: %s\n", u.Name, u.Age, u.Address, u.Phone, u.Email, u.Department)
}
func (u Patient) printInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, Phone: %s, Email: %s, MedicalRecord: %s\n", u.Name, u.Age, u.Address, u.Phone, u.Email, u.MedicalRecord)
}
func (u Driver) printInfo() {
	fmt.Printf("Name: %s, Age: %d, Address: %s, Phone: %s, Email: %s, DrivingLicense: %s, ExpierenceYears: %d\n", u.Name, u.Age, u.Address, u.Phone, u.Email, u.DrivingLicense, u.ExpierenceYears)
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

type Driver struct {
	User
	DrivingLicense  string
	ExpierenceYears int
}

type Salary interface {
	calculateSalary() float64
	paySalary()
}

func (s Staff) calculateSalary() float64 {

	return 50000.0
}
func (s Admin) calculateSalary() float64 {

	return 50000.0
}
func (s Driver) calculateSalary() float64 {

	return 50000.0
}

func (s Staff) paySalary() {
	fmt.Printf("Paying salary to Staff %s\n", s.Name)
}
func (s Admin) paySalary() {
	fmt.Printf("Paying salary to Admin %s\n", s.Name)
}
func (s Driver) paySalary() {
	fmt.Printf("Paying salary to Driver %s\n", s.Name)
}



func salaryDisbursement(s Salary) {
	s.paySalary()
	fmt.Printf("Salary Amount: %.2f\n", s.calculateSalary())
}



func main() {
	fmt.Println("Dependency Injection in Go")
	admin := Admin{
		User: User{
			Name:    "Alice",
			Age:     35,
			Address: "123 Admin St",
			Phone:   "123-456-7890",
			Email:   "admin@example.com",
		},
		AdminLevel: 1,
	}

	staff := Staff{
		User: User{
			Name:    "Bob",
			Age:     30,
			Address: "456 Staff Ave",
			Phone:   "987-654-3210",
			Email:   "staff@example.com",
		},
		Department: "Cardiology",
	}

	patient := Patient{
		User: User{
			Name:    "Charlie",
			Age:     25,
			Address: "789 Patient Rd",
			Phone:   "555-555-5555",
			Email:   "patient@example.com",
		},
		MedicalRecord: "Healthy",
	}

	driver := Driver{
		User: User{
			Name:    "David",
			Age:     40,
			Address: "101 Driver Ln",
			Phone:   "111-222-3333",
			Email:   "driver@example.com",
		},
		DrivingLicense:  "DL123456",
		ExpierenceYears: 15,
	}

	admin.printInfo()
	staff.printInfo()

	patient.printInfo()
	fmt.Println(driver.User.Address)
	patient.submitFeedback()
	staff.submitFeedback()
	driver.printInfo()

	salaryDisbursement(admin)
}