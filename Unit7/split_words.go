package main

import (
	"fmt"
	"strings"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	sentence := "this is the words"
	go sender(c1, sentence)
	go filter(c1, c2)
	printer2(c2)
}

func sender(downstream chan string, sentence string) {
	for _, v := range strings.Fields(sentence) {
		downstream <- v
	}
	close(downstream)
}

func filter(upstream, downstream chan string) {
	previous := ""
	for item := range upstream {
		if item != previous {
			downstream <- item
		}
		previous = item
	}
	close(downstream)
}

func printer2(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}
