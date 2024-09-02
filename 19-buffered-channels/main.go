package main

import (
	"fmt"
	"sync"
	"time"
)

func print(i int, wg *sync.WaitGroup) {
	fmt.Println("start", i)
	time.Sleep(2 * time.Second)
	fmt.Println("end", i)
	wg.Done()
}

func main() {

	ch := make(chan int, 3)
	ch <- 10
	ch <- 100
	ch <- 1000

	fmt.Println("len", len(ch))
	fmt.Println("cap", cap(ch))
	fmt.Println("read ", <-ch)
	fmt.Println("len", len(ch))
	fmt.Println("cap", cap(ch))

	// WaitGroup
	number := 3
	var wg sync.WaitGroup

	for i := 0; i < number; i++ {
		wg.Add(1)
		go print(i, &wg)
	}

	wg.Wait()
	fmt.Println("finished")
}
