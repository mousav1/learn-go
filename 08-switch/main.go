package main

import "fmt"

func main() {

	// switch

	a := 5
	switch a {
	case 1:
		fmt.Println(a)

	case 2:
		fmt.Println(a)

	case 3:
		fmt.Println(a)

	case 4:
		fmt.Println(a)

	// Duplicate cases are not allowed
	// case 4:
	// 	fmt.Println(num)
	case 6:
		fmt.Println(a)
	default:
		fmt.Println(5)
	}

	// Multiple expressions in case
	number := 5
	switch number {
	case 1, 2, 3, 4, 6:
		fmt.Println(number)

	default:
		fmt.Println("number false")
	}

	// Expressionless switch
	num := 10
	var code int
	switch {
	case num < 5:
		code = 10
	case num > 5 && num < 10:
		code = 20
	default:
		code = 30
	}
	fmt.Println(code)
}
