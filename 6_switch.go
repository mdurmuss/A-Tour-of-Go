package main

import (
	"fmt"
	"runtime"
)

func main() {

	// A switch statement is a shorter way to write a sequence of if - else statements.
	// It runs the first case whose value is equal to the condition expression.
	fmt.Print("Go runs on ")
	// os variable is the key.
	// no need to add break keyword to everycase.
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// var systemType = runtime.GOOS
	// fmt.Println(systemType)
}
