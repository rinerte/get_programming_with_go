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
		
		key:= keyword[index%len(keyword)]
		difference := key - 65
		result:=(char-rune(difference)-'A'+26)%26+'A'
		fmt.Printf("%c", result)
	}
}