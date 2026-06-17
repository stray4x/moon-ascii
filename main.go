package main

import (
	"fmt"
	"math"
)

func main() {
	drawCircle()
}

func drawCircle() {
	r := 15
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			dist := math.Sqrt(float64(x*x + y*y))
			if dist <= float64(r) {
				fmt.Print(". ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
