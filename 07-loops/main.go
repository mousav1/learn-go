package main

import "fmt"

func main() {
	// for loop syntax
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// break
	for i := 1; i < 5; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	// continue
	for i := 1; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("hello")
	}
}
