package main

import "fmt"

func main() {
	price, n := 100, 2
	c1, c2 := sum(price, n)
	// Blank Identifier
	c3, _ := sum(price, n)
	fmt.Println(c1, c2)
	fmt.Println(c3)
}

// Function declaration

// func functionname(parametername type) returntype {
//function body
//    }

// Sample Function
func sum(price, n int) (int, int) {
	a := price * n
	b := price + n
	return a, b
}
