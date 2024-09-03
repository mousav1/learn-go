package main

import (
	"fmt"
	"os"
)

func main() {
	// read-files
	contents, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	for _, v := range contents {
		fmt.Printf("%c \n", v)
	}

	// write-files
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("mohammad")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
