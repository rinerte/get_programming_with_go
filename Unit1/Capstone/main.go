package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var distance = 62100000
		
	fmt.Printf("%-20s %-6s %-15s %6s\n", "Spaceline", "Days", "Trip type", "Price")
    fmt.Println("--------------------------------------------------")

	for i:=0;i<10;i++{
		speed:=rand.Intn(30-16)+16
		tripType:=rand.Intn(2)+1
		var tripTypeCaption string
		if tripType==1{
			tripTypeCaption = "One-way"
		} else {
			tripTypeCaption = "Round-trip"
		}
		fmt.Printf("%-20s %-6d %-15s %-1s%4d\n", getRandSpaceline(), (distance/(speed*86400))*tripType, tripTypeCaption, "$", (speed+20)*tripType)
	}
}

func getRandSpaceline() string {
	switch r:=rand.Intn(3); r{
		case 0:
			return "Space Adventures"
		case 1:
			return "SpaceX"
		default: 
			return "Virgin Galactic"
	}
}