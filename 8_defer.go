package main

import "fmt"

func main() {
	//defer fmt.Println("World")

	//fmt.Printf("Hello ")

	// multi-defer
	fmt.Println("Counting")
	count := 0
	for i := 4; i > count; i-- {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
