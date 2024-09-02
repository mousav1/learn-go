package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "server1"
}

func server2(ch chan string) {
	ch <- "server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// ch := make(chan string)
	var ch chan string

	select {
	case s1 := <-ch:
		fmt.Println(s1)
	default:
		fmt.Println("no")
	}
}
