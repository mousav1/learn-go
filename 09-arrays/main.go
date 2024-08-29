package main

import "fmt"

func main() {
	// Arrays
	var a [3]int
	a[0] = 12
	a[1] = 22
	a[2] = 30
	fmt.Println(a)

	b := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println(b)

	c := [...]int{10, 20, 30, 40, 50, 60}
	d := c
	d[0] = 100
	fmt.Println(c)
	fmt.Println(d)

	e := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println("before", e)
	change(e)
	fmt.Println("after", e)

	f := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println("lenght", len(f))

	g := [...]int{10, 20, 30, 40, 50, 60}
	for i := 0; i < len(g); i++ {
		fmt.Println("element ", g[i])
	}
	for i, v := range g {
		fmt.Println(i, v)
	}

	for _, v := range g {
		fmt.Println(v)
	}

	h := [3][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	fmt.Println(h)

}

func change(number [6]int) {
	number[0] = 200
	fmt.Println("in function", number)
}
