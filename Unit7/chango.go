package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go printer(c)

	for i := 0; i < 10; i++ {
		c <- i
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		c <- i
	}
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	select {}
}

func printer(c chan int) {
	for {
		value := <-c
		fmt.Println(value)
	}
}
