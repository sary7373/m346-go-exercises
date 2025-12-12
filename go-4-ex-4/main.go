package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(C float64) float64 {
	return C*9/5 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(F float64) float64 {
	return (F - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println(convertCelsiusToFahrenheit(180))  //356
	fmt.Println(convertCelsiusToFahrenheit(30.8)) //87.44
	fmt.Println(convertCelsiusToFahrenheit(16.1)) //60.980000000000004
	fmt.Println(convertCelsiusToFahrenheit(0))    //32
	fmt.Println(convertCelsiusToFahrenheit(-6.7)) //19.939999999999998

	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println(convertFahrenheitToCelsius(300))   //148.88888888888889
	fmt.Println(convertFahrenheitToCelsius(67))    //19.444444444444443
	fmt.Println(convertFahrenheitToCelsius(32))    //0
	fmt.Println(convertFahrenheitToCelsius(0))     //-17.77777777777778
	fmt.Println(convertFahrenheitToCelsius(-10.5)) //-23.61111111111111
}
