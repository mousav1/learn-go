package main

import "fmt"

type User struct {
	name string
	age  int
	Address
}

type Address struct {
	code int
}

// Sample Method
func (a Address) printCode() {
	fmt.Println("code : ", a.code)
}

// Methods vs Functions
func (u User) print() {
	fmt.Println("name : ", u.name)
}
func print(u User) {
	fmt.Println("name : ", u.name)
}

func (u User) changeName(newName string) {
	u.name = newName
}

func (u *User) changeAge(newAge int) {
	u.age = newAge
}

func main() {
	user := User{
		name: "test",
		age:  30,
	}
	// Pointer Receivers vs Value Receivers
	fmt.Println(user)
	user.changeName("mohammad")
	fmt.Println(user)
	user.changeAge(35)
	fmt.Println(user)

	// Methods of anonymous struct fields
	user2 := User{
		name: "test",
		age:  30,
		Address: Address{
			code: 12,
		},
	}
	user2.printCode()
}
