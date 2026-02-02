package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	//user1 := User{Name:"Deepak" ,Age:20}
	addres1 := Address{HouseNo: 123, Street: "MG Road", City: "Bangalore", State: "Karnataka", Coutnty: "India", PinCode: 560001}
	//user1 := createNewUser("Deepak", 20, addres1)
	user1 := User{Name: "Deepak", Age: 20, UserAddress: addres1}
	//fmt.Println(user1.Name)
	//fmt.Println(user1.Age)
	fmt.Println(user1.name())
	fmt.Println(user1.age())
	//we get get a Struct of address from our User
	var parent Printer
	
	address := user1.address()
	parent = address

	parent.Print()
	// How we can attachment methods for Address Struct
	fmt.Println(address.houseNo())
	fmt.Println(address.street())
	fmt.Println(address.city())
	fmt.Println(address.state())
	fmt.Println(address.country())
	fmt.Println(address.pinCode())

	info, err := getUserInfo(user1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(info)
	}
}

/**
Create struct
Factory Convention
Public View
Composition
**/

type Address struct {
	HouseNo int
	Street  string
	City    string
	State   string
	Coutnty string
	PinCode int
}

type Printer interface{
	Print()
}

func (a Address) houseNo() int {

	return a.HouseNo
}

func (a Address) street() string {

	return a.Street
}
func (a Address) city() string {

	return a.City
}
func (a Address) state() string {
	return a.State
}
func (a Address) country() string {
	return a.Coutnty
}
func (a Address) Print() {
	fmt.Println(fmt.Sprintf("House No :	%d",a.HouseNo))
}
func (a Address) pinCode() int {

	return a.PinCode
}

/*type User struct {
	Name string
	Age  int

}*/

func (u User) name() string {

	return u.Name
}

func (u User) age() int {

	return u.Age
}

func (u User) address() Address {

	return u.UserAddress
}

// func createNewUser(name string, age int) User {

// 	return User{Name: name, Age: age}

// }

//Public View
// func getUserInfo(user User) (string, error) {

// 	if user.Name == "" {
// 		return "", fmt.Errorf("Name field can't be empty")
// 	} else if user.Age == 0 {
// 		return "", fmt.Errorf("Age value can't be zero:%d", user.Age)

// 	}
// 	return fmt.Sprintf("User Name : %s, User Age: %d", user.Name, user.Age), nil

// }

type User struct {
	Name        string
	Age         int
	UserAddress Address
}

func createNewUser(name string, age int, address Address) User {

	return User{Name: name, Age: age, UserAddress: address}

}

//Public View
func getUserInfo(user User) (string, error) {

	if user.Name == "" {
		return "", fmt.Errorf("Name field can't be empty")
	} else if user.Age == 0 {
		return "", fmt.Errorf("Age value can't be zero:%d", user.Age)

	} else if user.UserAddress == (Address{}) {
		return "", fmt.Errorf("Address value is required:%s", user.UserAddress)
	}
	return fmt.Sprintf("User Name : %s, User Age: %d, User Address: %s", user.Name, user.Age, user.UserAddress), nil

}

//Methods with Struct
