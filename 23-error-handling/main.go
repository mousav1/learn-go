package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../test.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println(pErr.Path)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name())
}
