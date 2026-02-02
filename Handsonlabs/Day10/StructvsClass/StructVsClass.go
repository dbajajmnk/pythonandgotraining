package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	//user1 := User{Name:"Deepak" ,Age:20}
	addres1 := Address{HouseNo:123,Street:"MG Road",City:"Bangalore",State:"Karnataka",Coutnty:"India",PinCode:560001}	
	//user1 := createNewUser("Deepak", 20, addres1)
	user1 := User{Name:"Deepak", Age:20,UserAddress:addres1}
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)
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
	Street string
	City string
	State string
	Coutnty string
	PinCode int
	
}

/*type User struct {
	Name string
	Age  int
	
}



func createNewUser(name string, age int) User {

	return User{Name: name, Age: age}

}

//Public View
func getUserInfo(user User) (string, error) {

	if user.Name == "" {
		return "", fmt.Errorf("Name field can't be empty")
	} else if user.Age == 0 {
		return "", fmt.Errorf("Age value can't be zero:%d", user.Age)

	}
	return fmt.Sprintf("User Name : %s, User Age: %d", user.Name, user.Age), nil

}
*/

type User struct {
	Name string
	Age  int
	UserAddress Address
}



func createNewUser(name string, age int, address Address) User {

	return User{Name: name, Age: age,UserAddress:address}

}

//Public View
func getUserInfo(user User) (string, error) {

	if user.Name == "" {
		return "", fmt.Errorf("Name field can't be empty")
	} else if user.Age == 0 {
		return "", fmt.Errorf("Age value can't be zero:%d", user.Age)

	} else if  user.UserAddress == (Address{}) {
		return "", fmt.Errorf("Address value is required:%s", user.UserAddress)
	}
	return fmt.Sprintf("User Name : %s, User Age: %d, User Address: %s", user.Name, user.Age, user.UserAddress), nil

}