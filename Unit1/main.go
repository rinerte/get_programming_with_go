package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	rand.Seed(time.Now().UnixNano())
	
	for i:=0;i<10;i++{
		year := rand.Intn(2025)+1
		month := rand.Intn(2) + 1
		daysInMonth := 31
		switch month {
			case 2:
				daysInMonth = 28
				if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
					daysInMonth = 29
				}			
			case 4, 6, 9, 11:
				daysInMonth = 30
		}
		
			day := rand.Intn(daysInMonth) + 1
			fmt.Println(era, year, month, day)
	}	
}