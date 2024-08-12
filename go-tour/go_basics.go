package main

import (
	"fmt"
	"math/rand"
)

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
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
