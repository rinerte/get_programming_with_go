package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64


func kelvinToCelsius(k kelvin) celsius{
	return celsius(k-273.15)
}
func celsiusToFahrenheit(c celsius) fahrenheit{
	return fahrenheit(c*9.0/5.0 + 32.0)
}
func kelvinToFahrenheit(k kelvin) fahrenheit{
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func celsiusToKelvin(c celsius) kelvin{
	return kelvin(c+273.15)
}

// method
func (k kelvin) celsius() celsius{
	return celsius(k-273.15)
}
func main(){

	//var gradus fahrenheit = 2.0
	fmt.Println(kelvin(273.0).celsius())
	fmt.Println(kelvinToFahrenheit(0.0))

	fmt.Println(celsiusToKelvin(127))
	
	
}