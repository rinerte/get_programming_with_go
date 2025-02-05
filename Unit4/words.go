package main

import (
	"fmt"
	"strings"
)

func main(){
	count:= map[string]int{}

	text:=`As far as eye could reach he saw nothing but the stems of the great plants about him
	receding in the violet shade, and far overhead the multiple transparency of huge leaves
	filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
	he felt able he ran again; the ground continued soft and springy, covered with the same
	resilient weed which was the first thing his hands had touched in Malacandra. Once or
	twice a small red creature scuttled across his path, but otherwise there seemed to be no
	life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
	and alone in a forest of unknown vegetation thousands or millions of miles beyond the
	reach or knowledge of man.`

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\'' 
	})

	for _, word := range words {
		word = strings.Trim(word, ".,!-—")  
		word = strings.ToLower(word)
		if word != "" { 
			count[word]++
		}
	}

	for word,number:= range count{
		if number>=2 {
			fmt.Printf("%s : %d times\n",word,number)
		}
	}

	fmt.Println(count["fear—except"])
}