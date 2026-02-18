package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   *int   `json:"age"`
	Admin *bool  `json:"admin"`
}

func basicExample() {

	fmt.Println("=== Basic Example ===")

	p := new(42)

	fmt.Println("value:", *p)
	fmt.Println("type:", fmt.Sprintf("%T", p))
}

func structExample() {

	fmt.Println("\n=== Struct Example ===")

	user := User{
		Name:  "Ali",
		Age:   new(30),
		Admin: new(true),
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")

	fmt.Println(string(jsonData))
}

func sliceExample() {

	fmt.Println("\n=== Slice Example ===")

	nums := new([]int{1, 2, 3, 4})

	fmt.Println("slice:", *nums)
}

func functionExample() {

	fmt.Println("\n=== Function Example ===")

	getName := func() string {
		return "Mohammad"
	}

	namePtr := new(getName())

	fmt.Println("name:", *namePtr)
}

func main() {

	basicExample()

	structExample()

	sliceExample()

	functionExample()

}
