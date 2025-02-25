package main

import (
	"fmt"
)

func main(){

	bolden:="string 1"

	pointer:=&bolden

	fmt.Println(*pointer)

	bolden="string 2"

	fmt.Println(*pointer)
}