package main

import(
	"fmt"
	"math/rand"
)

func main(){
	var number = 25
	var guess = -1

	var top = 100;
	var bot = 0;
	for number!=guess{

		guess = rand.Intn(top-bot)+bot
		fmt.Printf("I guess it is %v \n",guess)

		if(guess==number){
			fmt.Println("Yes, It Is!")
		} else if(guess<number){
			bot = guess+1
			fmt.Println("No, it is not :(")
		} else {
			top = guess-1
			fmt.Println("No, it is not :(")
		}
	}
}