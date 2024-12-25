package main

import(
	"fmt"
)

func main(){
	var command = "asdfa2sdfasdf"

	switch command{
	case "go east":
		fmt.Println("Go east")
	case "go inside":
		fmt.Println("Go inside")
		fallthrough
	case "asdfasdfasdf":
		fmt.Println("This is fallthroudh")
	default:
		fmt.Println("This is default")
	}
}