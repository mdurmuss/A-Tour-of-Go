package main

import (
	"fmt"
)

func main() {

	// For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While
	var sum = 15
	for sum < 20 {
		fmt.Println(sum)
		sum++
	}

	// Infinite loop and If Condition
	var x = 0
	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}

	// I stay at for loop basics!

}
