package main

import "fmt"

func kelvinToCelsius(kelvin float64) float64{
	return kelvin-273.15
}
func celsiusToFahrenheit(celsius float64) float64{
	return celsius*9.0/5.0 + 32.0
}
func kelvinToFahrenheit(kelvin float64) float64{
	return celsiusToFahrenheit(kelvinToCelsius(kelvin))
}
func main(){
	fmt.Println(kelvinToFahrenheit(0.0))
}