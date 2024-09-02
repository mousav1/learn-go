package main

import "fmt"

func hello(data chan bool) {
	fmt.Println("hello1")
	data <- true
}

func cal1(number int, ch1 chan int) {
	number *= 2
	ch1 <- number
}

func cal2(number int, ch2 chan int) {
	number += 2
	ch2 <- number
}

func sendData(data chan<- int) {
	data <- 10
}

func getNumers(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}
func main() {

	// Declaring channels
	var a chan int
	fmt.Println(a)

	b := make(chan int)
	fmt.Printf("%T", b)

	// Sending and receiving from a channel
	data := <-a
	a <- data
	fmt.Println(data)

	ch1 := make(chan bool)
	go hello(ch1)
	<-ch1
	fmt.Println("main2")

	number := 10
	ch2 := make(chan int)
	ch3 := make(chan int)
	go cal1(number, ch2)
	go cal2(number, ch3)
	s1, s2 := <-ch2, <-ch3
	fmt.Println(s1, s2)

	// Unidirectional channels
	// ch4 := make(chan<- int)
	ch4 := make(chan int)
	go sendData(ch4)
	f := <-ch4
	fmt.Println(f)

	// Closing channels
	ch5 := make(chan int)
	go getNumers(ch5)
	for v := range ch5 {
		fmt.Println(v)
	}
}
