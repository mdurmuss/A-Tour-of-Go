package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Friday?")
	today := time.Now().Weekday()

	//	fmt.Printf("%s\n", (today + 1))

	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 5:
		fmt.Println("In five days.")
	default:
		fmt.Println("Too far away.")
	}
}
