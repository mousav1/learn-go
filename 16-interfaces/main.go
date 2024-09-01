package main

import "fmt"

type Calculator interface {
	getSpeed() int
}

type car struct {
	name  string
	speed int
}

func (c car) getSpeed() int {
	return c.speed
}

type User struct {
	name  string
	speed int
}

func (c User) getSpeed() int {
	return c.speed
}

func Print(c Calculator) {
	fmt.Println(c)
}

// Empty interface
func newPrint1(i interface{}) {
	fmt.Println(i)
}

// Type assertion
func newPrint2(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	// Declaring and implementing an interface
	car1 := car{
		name:  "BMW",
		speed: 2000,
	}
	speed := car1.getSpeed()
	fmt.Println(speed)
	Print(car1)

	user1 := User{
		name:  "test",
		speed: 2000,
	}
	// Print(user1)
	newPrint1(car1)
	newPrint1(user1)

	a := 21
	newPrint2(a)
	b := "asd"
	newPrint2(b)

}
