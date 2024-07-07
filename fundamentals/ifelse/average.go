package main

import "fmt"

func nota(n float64) string {
	switch {
	case n < 3:
		return "D"
	case n < 5:
		return "C"
	case n < 7:
		return "B"
	case n <= 10:
		return "A"
	default:
		return "E"
	}
}

func main() {
	resultado := nota(0)
	fmt.Println("A nota é:", resultado)
}
