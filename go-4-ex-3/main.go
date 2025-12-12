package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	D := math.Pow(b, 2) - 4*a*c
	if D < 0 {
		return []float64{}
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		if D == 0 {
			return []float64{x1}
		} else {
			x2 := (-b - math.Sqrt(D)) / (2 * a)
			return []float64{x1, x2}
		}
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1)) //-0.3333333333333333, -1
	fmt.Println(computeQuadraticFormula(2, 4, 2)) //-1
	fmt.Println(computeQuadraticFormula(3, 4, 2)) //leer
}
