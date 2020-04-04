package main

// libraries
import (
	"fmt"
	"math"
	"math/rand"
)

// first function
func add(x, y int) int {
	var z int
	z = x + y
	return z
}

// swap between 2 variables
// (string, string) --> return types!
func swapper(x, y string) (string, string) {

	return y, x
}

// A return statement without arguments returns the
// named return values. This is known as a "naked" return.
func nakedReturn(sum int) (x, y int) {
	x = sum / 9 * 4
	y = sum - x
	return
}

// defining variables
func defineVariable() {
	var c, python bool
	python = true
	fmt.Println(c, python)
}

// defining variables
func defineVariable2() {
	var i, j int = 1, 2
	var c, python, java = true, false, "NO!"
	k := 3

	fmt.Println("i:", i, "j:", j, "c,python,java:", c, python, java)
	fmt.Println(k)
}

func main() {
	fmt.Println("My favorite number is : ", rand.Float32())

	fmt.Println("My first go square process: ", math.Sqrt(25))

	fmt.Println("Pi is : ", math.Pi)

	fmt.Println(add(15, 25))

	a, b := swapper("hello", "world")

	fmt.Println("Swap: ", a, b)

	fmt.Println(nakedReturn(9))

	defineVariable()

	defineVariable2()

	// I finished here!
	// https://tour.golang.org/basics/11
}
