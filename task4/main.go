package main

import (
	"fmt"
	"math"
)

func main() {
	xValues := []float64{0.2, 0.3, 0.38, 0.43, 0.57}

	fmt.Println("x\t\ty")
	fmt.Println("------------------------")

	for i, x := range xValues {
		y := (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
		fmt.Printf("x%d = %.2f\t%.6f\n", i+1, x, y)
	}
}
