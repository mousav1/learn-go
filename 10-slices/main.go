package main

import "fmt"

func main() {
	// Creating a slice
	a := [6]int{80, 78, 56, 55, 23, 10}
	var b []int = a[1:4]
	fmt.Println(b)

	c := []int{80, 78, 56, 55, 23, 10}
	fmt.Println(c)

	// modifying a slice
	d := [6]int{80, 78, 56, 55, 23, 10}
	var e []int = d[1:4]
	fmt.Println("before ", d)
	for i := range e {
		e[i]++
	}
	fmt.Println("after ", e)

	// length and capacity of a slice
	f := [6]int{80, 78, 56, 55, 23, 10}
	var g []int = f[1:4]
	fmt.Println(len(g), cap(g))

	// creating a slice using make
	h := make([]int, 3, 5)
	fmt.Println(h)

	// Appending to a slice
	i := []string{"a", "b", "c"}
	i = append(i, "d")
	fmt.Println(i)

	// Passing a slice to a function
	j := []int{80, 78, 56, 55, 23, 10}
	fmt.Println("befor ", j)
	change(j)
	fmt.Println("after ", j)

}

func change(numbers []int) {
	for i := range numbers {
		numbers[i] -= 1
	}
}
