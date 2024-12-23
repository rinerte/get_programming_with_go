package main

import(
	"fmt"
)

func main(){
	var year = 2104

	var leap = (year%4==0 && year%100!=0) || year%400 == 0

	if(leap){
		fmt.Println("LEAP")
	} else {
		fmt.Println("no leap")
	}
}