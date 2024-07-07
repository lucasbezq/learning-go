package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2

	area := math.Pow(radius, 2)
	fmt.Println("The area of the circunference is ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
}
