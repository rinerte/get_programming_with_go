package main

import "fmt"

func main(){

	slice:= []int{1}

	for i:=0;i<100; i++{
		slice = append(slice, i)
		fmt.Println(cap(slice))
	}

}