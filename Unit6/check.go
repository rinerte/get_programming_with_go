package main

import(
	"fmt"
)

type some struct{
	name string
	too string
}

func main(){
	b:=some{"Name","boo"}

	fmt.Printf("%v\n",b)
}