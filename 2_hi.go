package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14

func main() {

	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// DEFAULT VALUES
	var i int
	var f1 float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f1, b, s)

	// Variable types
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = f
	fmt.Println(x, y, z)

	// Constant Variables
	const World = "世界"
	fmt.Printf("Type: %T Value: %v \n", pi, pi)
	fmt.Println("Hello", World)
}
