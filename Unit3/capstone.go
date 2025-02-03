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
	
	const (
		line = "======================="
		rowFormat = "| %8s | %8s |\n"
		numberFormat = "%.1f"
	)

	type rowDataFunction func(value float64) (string,string)
	type StrStr func() (string,string)
	type FlFl func() (float64, float64)

	func drawTable(h1,h2 string, min, max float64, typePairFunction rowDataFunction){
		fmt.Println(line)
		fmt.Printf(rowFormat, h1,h2)
		fmt.Println(line)

		for value:=min; value<=max; value+=5.0 {
			var c1,c2 = typePairFunction(value)
			fmt.Printf(rowFormat,c1,c2)
		}
		fmt.Println(line)
	}

	func celsiusToFahrenheitPair(value float64) (string, string) {
		return fmt.Sprintf(numberFormat,value), fmt.Sprintf(numberFormat,celsius(value).fahrenheit())
	}
	func fahrenheitToCelsiusPair(value float64) (string, string) {
		return fmt.Sprintf(numberFormat,value), fmt.Sprintf(numberFormat,fahrenheit(value).celsius())
	}

func main(){
	drawTable("ºC","ºF",-40.0,20.0,celsiusToFahrenheitPair)
	drawTable("ºF","ºC",-40.0,20.0,fahrenheitToCelsiusPair)
}