package main

import "fmt"

func main() {
	// variadic function
	numbers := []int{90, 10, 900, 60}
	find(10, numbers...)
}

func find(number int, numbers ...int) {
	found := false
	for i, v := range numbers {
		if v == number {
			fmt.Println(number, "found at index", i)
			found = true
		}
	}

	if !found {
		fmt.Println(number, "not found")
	}
}
