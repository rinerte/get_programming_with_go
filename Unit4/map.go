package main

import (
	"fmt"
	"strconv"
)

func main(){

	//map[keyType]valueType
	 mymap:= make(map[string]string,0)

	fmt.Println("FIrst = "+ mymap["First"] +" "+strconv.Itoa(len(mymap)))
	mymap["First"] = "First"
	fmt.Println("FIrst = "+ mymap["First"])
}