package main

import "fmt"

func main() {
	// Declaring pointers
	b := 123
	var a *int = &b
	fmt.Println(a)

	// pointers
	c := 123
	var d *int
	if d == nil {
		fmt.Println(d)
		d = &c
		fmt.Println(d)
	}

	// Creating pointers using the new function
	e := new(int)
	fmt.Println(e)

	// Dereferencing a pointer
	f := 123
	g := &f
	fmt.Println(g)
	fmt.Println(*g)

	// Passing pointer to a function
	h := 123
	fmt.Println(h)
	i := &h
	change(i)
	fmt.Println(h)
}

// Passing pointer to a function
func change(val *int) {
	*val = 456
}

// Returning pointer from a function
func change2(val *int) *int {
	return val
}
