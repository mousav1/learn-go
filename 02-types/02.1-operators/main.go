package main

import "fmt"

func main() {
	// x + y
	// x - y
	// x * y
	// x / y
	// x % y

	a := -88
	b := 85
	c := 15
	fmt.Println(a + b)
	fmt.Println(c + b)
	fmt.Println(c * b)
	fmt.Println(b % c)

	c += 10 //=> c = c + 10
	c -= 10 //=> c = c - 10
	c *= 10 //=> c = c * 10
	c /= 10 //=> c = c / 10
	c %= 10 //=> c = c % 10

}
