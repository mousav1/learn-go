package main

import "fmt"

func main() {
	// // Declaring variables
	var a int
	var b string
	fmt.Println(a)
	fmt.Println(b)

	// Initializing Variables
	a = 10 //assignment
	fmt.Println(a)

	// Creating and initializing variables
	var c = 10
	var d = "topdemy.ir"
	fmt.Println(c)
	fmt.Println(d)

	// Declaring multiple variables
	var e, f int
	fmt.Println(e, f)
	e = 10
	f = 100
	fmt.Println(e, f)
	// Creating and initializing multiple variables
	var g, h = 10, 100
	fmt.Println(g, h)

	// Variable Declaration Block
	var (
		i = 10
		j = "mohammad"
		k = 1000
	)

	fmt.Println(i, j, k)

	//Short variable declaration
	var l int
	var m int
	l, m, n := 10, 20, 30
	fmt.Println(l, m, n)
}
