package main

import (
	"fmt"
)

func main(){
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	//ABCZ
	//var difference = x - 'A'
	
	for index, char:=range cipherText {		
		key:= rune(keyword[index%len(keyword)])
		fmt.Printf("%c", (char-key+26)%26+'A')
	}
}