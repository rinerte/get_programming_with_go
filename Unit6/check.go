package main

import(
	"fmt"
)

type some struct{
	name string
	too string
}

func (a *some) rename(b string){
	a.name = b
}

func main(){
	a:=some{"oldName","ha"}
	fmt.Printf("%v\n",a.name)
	a.rename("newName")
	fmt.Printf("%v\n",a.name)
}