package main

import "fmt"

func print() {
	fmt.Println("mohammad")
}
func printName() {
	fmt.Println("mousavi")
}
func value(a int) {
	fmt.Println(a)

}
func main() {
	defer func() { fmt.Println("mohammad") }()
	fmt.Println("hello")

	a := 1
	defer value(a)
	a = 10
	fmt.Println(a)

	defer printName()
	fmt.Println("hello")
}
