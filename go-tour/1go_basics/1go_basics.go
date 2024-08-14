package main

import (
	"fmt"
	"math/rand"
)

// var in decleares a list of variables with type in the end
var c, python, java int = 1, 2, 3

// basic go types, but there are types such as rune=int32 and complex.
// variables declared without any values will default to zero values. "", false 0
// classic type conversion like float64(i)
var (
	boogie bool    = false
	number int     = 1
	streng string  = "hei"
	komma  float32 = 1.1
)

// Pretty much same as var but can not be declared with := and are fixed
const con = 10

// using import
func rando() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// func concept
func add(x, y int) int {
	return x + y
}

// Func multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println((add(42, 113)))
	// := init variable and type, this only works inside a function?
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(c, python, java)
	fmt.Println(boogie, number, streng, komma)
}
