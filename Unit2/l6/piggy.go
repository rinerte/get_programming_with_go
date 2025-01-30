package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var total int

	for total < 2000 {
		switch coin:=rand.Intn(3); coin{
		case 0:
			total+=5
			fmt.Println(" 0.05$ added")
		case 1:
			total+=10
			fmt.Println(" 0.10$ added")
		case 2:
			total+=25
			fmt.Println(" 0.25$ added")
		}
	}
	
	fmt.Printf("%25s %05.2f\n","In the piggy bank now", float64(total)/100.0)

}
