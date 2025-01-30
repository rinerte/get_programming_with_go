package main

import (
	"fmt"
)

func main() {
	
	code:="L fdph, L vdz, L frqtxhuhg."

	for _, s:=range code{
		if((s<='z' && s>='a') || (s<='Z' && s>='A')){
			s-=3
		}
		fmt.Printf("%c",s)
	}

}
