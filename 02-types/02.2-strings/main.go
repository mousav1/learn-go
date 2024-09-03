package main

import "fmt"

func printCh(s string) {
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	fmt.Print("hi", "mohmmad")
	fmt.Print("aaa")
	fmt.Println("hi", "mohmmad")
	fmt.Println("aaa")
	fmt.Printf("hi %s test", "mohmmad")
	// fmt.Printf("hi %s %d %T %f %v test",)

	var a string = "test"
	b := "test2"

	fmt.Println(a == b)

	c := a + " " + b

	fmt.Println(c)

	printCh(a)

	d := []byte{121, 34, 213, 65, 10}
	str := string(d)
	fmt.Println(str)
}
