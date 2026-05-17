package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		xn   = 0.11
		xk   = 0.36
		step = 0.05
	)

	fmt.Println("x\t\ty")
	fmt.Println("------------------------")

	for x := xn; x <= xk+1e-9; x += step {
		y := (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
		fmt.Printf("%.2f\t\t%.6f\n", x, y)
	}
}
