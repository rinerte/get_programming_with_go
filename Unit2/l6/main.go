package main

import(
	"fmt"
	//"math"
)

func main(){
	// var f32 float32 = math.Pi
	// f64 := math.Pi

	// fmt.Println(f32)
	// fmt.Println(f64)

	// fmt.Printf("%05.2f\n", f64)

	//var piggy float64

	// for i:=0;i<11;i++{
	// 	piggy+=0.1
	// }

	// fmt.Println(piggy)

	//var i int64   =  4.03e8
	//fmt.Println(i)
	// var j = 1
	// var k int64 = 1

	// fmt.Printf("%T for %[1]v\n",i)
	// fmt.Printf("%T for %[1]v\n",j)
	// fmt.Printf("%T for %[1]v\n",k)

// 	const t = 1000000000000000000000000000000
// 	const y = 1000000000000000000
// 	fmt.Println(t/y)
// 	fmt.Printf("%T\n",t/y)


// 	fmt.Println(`awdaw
// 	awdawd
// awdaw  awdawd
// 	awdawd`)

		//  str:="lowercase string UPPERCASE STRING"
		 
		//  for i:=0;i<len(str);i++{
		// 	fmt.Printf("%c",str[i]-'a'+'A')
		//  }

		// fmt.Printf("%c\n",c)

		ss := "English Противень  и ячмень"

		for i, s:= range ss {
			fmt.Printf("%v %c\n", i, s)
		}
}
