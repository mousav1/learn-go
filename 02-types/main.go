package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
	   . bool
	   . string
	   . Numeric Types
	   	uint8       8 bit unsigned integers (0 to 255)
	   	uint16      16 bit unsigned integers (0 to 65535)
	   	uint32      32 bit unsigned integers (0 to 4294967295)
	   	uint64      64 bit unsigned integers (0 to 18446744073709551615)
	   	uint 	    32 or 64 bit unsigned integers depending on the underlying architecture

	   	int8        8 bit signed integers (-128 to 127)
	   	int16       16 bit signed integers integers (-32768 to 32767)
	   	int32       32 bit signed integers (-2147483648 to 2147483647)
	   	int64       64 bit signed integers (-9223372036854775808 to 9223372036854775807)
	   	int 	   represents 32 or 64 bit integers depending on the underlying architecture

	   	float32     32 bit floating point numbers
	   	float64     64 bit floating point numbers
	   	complex64   complex numbers with float32 real and imaginary parts
	   	complex128  complex numbers with float64 real and imaginary parts

	   	byte        alias for uint8
	   	rune        alias for int32
	*/

	// bool
	var a bool = true
	var b bool = false
	c := a || b
	d := a && b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// string
	var e string = "test1"
	var f string = "test2"
	g := e + " " + f
	fmt.Println(g)

	// int
	var i int = 32
	j := 90

	fmt.Println("value of a is ", i, "and b is ", j)
	fmt.Printf("type of a is %T, size of a is %d bytes", i, unsafe.Sizeof(i))

	// uint
	var k uint = 32
	l := 90

	fmt.Println("value of a is ", k, "and b is ", l)
	fmt.Printf("type of a is %T, size of a is %d bytes", k, unsafe.Sizeof(k))

	// float32
	var m, n float64 = 5.32, 1.12
	fmt.Println("value of a is ", m, "and b is ", n)
	fmt.Printf("type of a is %T, size of a is %d bytes", m, unsafe.Sizeof(m))

	//complex
	o := complex(5, 6)
	p := 8 + 33i
	sum := o + p
	fmt.Println("value of a is ", sum)

}
