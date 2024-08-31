package main

import "fmt"

func main() {

	// create a map
	a := make(map[string]int)
	fmt.Println(a)

	// Adding items to a map
	b := map[string]int{
		"name": 100,
		"test": 123,
	}
	b["asd"] = 123123
	fmt.Println(b)

	// nil map panics
	var c map[string]int
	c["asd"] = 123123
	fmt.Println(c)

	// Retrieving value for a key from a map
	d := map[string]int{
		"name": 100,
		"test": 123,
	}
	fmt.Println(d["name"])

	// Checking if a key exists
	if value, ok := d["name"]; ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println(ok)
	}

	// Deleting items from a map
	e := map[string]int{
		"name": 100,
		"test": 123,
	}
	fmt.Println("before", e)
	delete(a, "name")
	fmt.Println("after", e)

	// Iterate over all elements in a map
	f := map[string]int{
		"name": 100,
		"test": 123,
	}

	for i, v := range f {
		fmt.Println(i, v)
	}

	// Maps are reference types
	g := map[string]int{
		"name": 100,
		"test": 123,
	}

	test := g
	test["name"] = 99
	fmt.Println(g)
}
