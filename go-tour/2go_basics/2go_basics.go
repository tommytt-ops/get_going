package main

import (
	"fmt"
	"math"
)

// standard loop
func loopy() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func betterloopy() {
	sum := 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println(sum)
}

func evenbetterloopy() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteloop() {
	for {
		fmt.Print("hello")
	}
}

// Not correct math ik ;(
func ifstatement(x float64) float64 {
	if x < 0 {
		x = -1 * x
	}
	return math.Sqrt(x)
}

// Bad example
func shortifstatement(x float64) float64 {
	if i := x; i < 0 {
		x = -1 * x
	}
	return math.Sqrt(x)
}
func Sqrt(x float64) (float64, float64) {
	z := 1.0
	i := 0.0
	for i < 1000 {
		i += 1
		e := math.Round(z * 1000)
		z -= (z*z - x) / (2 * z)
		f := math.Round(z * 1000)
		if e == f {
			break
		}
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(2))
}
