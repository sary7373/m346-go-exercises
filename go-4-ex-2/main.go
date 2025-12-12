package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSites struct {
	a float64
	b float64
}

func Hypotenuse(sites ShortSites) float64 {
	return math.Sqrt(math.Pow(sites.a, 2) + math.Pow(sites.b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(3, 4))     //5
	fmt.Println(computeHypotenuse(17, 9))    //19.235384061671343
	fmt.Println(computeHypotenuse(1.6, 1.7)) //2.3345235059857505

	fmt.Println(Hypotenuse(ShortSites{3, 4}))     //5
	fmt.Println(Hypotenuse(ShortSites{17, 9}))    //19.235384061671343
	fmt.Println(Hypotenuse(ShortSites{1.6, 1.7})) //2.3345235059857505
}
