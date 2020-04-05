package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}
}

func main() {
	fmt.Println(sqrt(25), sqrt(-16))

	// usage of sprint
	var x = 10
	var test = fmt.Sprintln("x esittir :", x)
	fmt.Printf("%T\n", test)

}
