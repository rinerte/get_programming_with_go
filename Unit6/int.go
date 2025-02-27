package main

import "fmt"

type object interface{}
func main(){
	var s object
	s= 4
	fmt.Println(s)
	s = "Asdasd"
	fmt.Println(s)
}