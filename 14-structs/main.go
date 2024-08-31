package main

import (
	"fmt"
	"test/model"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Car struct {
	string
	int
}

func main() {
	// structs
	user1 := Person{
		lastName:  "mousavi",
		age:       100,
		firstName: "mohammad",
	}
	user2 := Person{
		"test",
		"test2",
		2,
	}
	fmt.Println(user1)
	fmt.Println(user2)

	// Creating anonymous structs
	user3 := struct {
		firstName string
		lastName  string
		age       int
	}{
		lastName:  "mousavi",
		age:       1,
		firstName: "mohammad",
	}
	fmt.Println(user3)

	// Accessing individual fields of a struct
	user4 := Person{
		lastName:  "mousavi",
		age:       1,
		firstName: "mohammad",
	}
	fmt.Println(user4.firstName)

	// Zero value of a struct
	var user5 Person
	user5 = Person{
		firstName: "aaa",
		lastName:  "bbb",
	}
	fmt.Println(user5.firstName)
	fmt.Println(user5.lastName)
	fmt.Println(user5.age)

	// Anonymous fields
	car1 := Car{
		string: "asdasd",
		int:    123,
	}
	fmt.Println(car1)

	// Exported structs and fields
	user6 := model.User{
		LastName:  "mousavi",
		FirstName: "mohammad",
		Address: model.Address{
			City: "asd",
		},
	}
	fmt.Println(user6.City)
}
