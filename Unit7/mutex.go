package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	go inc(&i)
	go inc(&i)
	go inc(&i)
	time.Sleep(3 * time.Second)
	fmt.Println(i)
}

func inc(n *int) {
	for i := 0; i < 100000000; i++ {
		*n++
	}
}
