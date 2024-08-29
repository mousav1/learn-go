package main

import "fmt"

func main() {

	// If statement syntax
	number := 10

	if number == 2 {
		fmt.Println("number 2")
	} else {
		fmt.Println("number :", number)
	}

	// If else statement
	num := 10
	var code int

	if num < 5 {
		code = 10
	} else if num > 5 && num < 10 {
		code = 20
	} else {
		code = 30
	}
	fmt.Println(code)

	// If with assignment
	var price int

	if age := 10; age < 5 {
		price = 10
	} else if age > 5 && age < 10 {
		price = 20
	} else {
		price = 30
	}

	fmt.Println(price)
}
