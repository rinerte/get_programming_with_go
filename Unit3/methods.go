package main

import "fmt"

	type celsius float64
	type fahrenheit float64
	type kelvin float64

	func (x kelvin) celsius() celsius {
		return celsius(x-273.15)
	}
	func (x kelvin) fahrenheit() fahrenheit{
		return x.celsius().fahrenheit()
	}
	func (x fahrenheit) celsius() celsius {
		return celsius((x-32.0) * 5.0/9.0)
	}
	func (x fahrenheit) kelvin() kelvin{
		return x.celsius().kelvin()
	}
	func (x celsius) fahrenheit() fahrenheit {
		return fahrenheit(x*9.0/5.0 + 32.0)
	}
	func (x celsius) kelvin() kelvin{
		return kelvin(x+273.15)
	}
	

func main(){
	var kGrad kelvin = 273.15

	fmt.Println("Kelvin to celsius = ", kGrad.celsius())
	fmt.Println("Kelvin to fahrenheit = ", kGrad.fahrenheit())
	fmt.Println("Kelvin to all and down to Kelvin = ", kGrad.celsius().fahrenheit().kelvin().fahrenheit().celsius().kelvin())


}