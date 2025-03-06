package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go sender(c1)
	go filter(c1, c2)
	printer2(c2)
}

func sender(downstream chan string) {
	for _, v := range []string{"aa", "aa", "bb", "bb", "bb", "cc", "cc", "cc"} {
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
