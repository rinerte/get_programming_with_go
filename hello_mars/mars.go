//comment
package main

import "fmt"

func main(){
	fmt.Print("My weight on Mars is ")
	fmt.Print(80.0*0.3783)
	fmt.Print(" kg. And I would be ")
	fmt.Print(30*365.2425 / 687)
	fmt.Print(" years old.")

	fmt.Printf("\nNumbers: %v %v %v %v \n New LIne\n",1,2,3,4)

	fmt.Println("Table:")
	fmt.Printf("%-20v |%5v\n","Name","Number")
	fmt.Printf("%-20v |%5v\n","John","1")
	fmt.Printf("%-20v |%5v\n","Not John","2")
	fmt.Printf("%-20v |%5v\n","Barack","3")
}