package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i := 0
	var mutex sync.Mutex

	go inc(&i, &mutex)
	go inc(&i, &mutex)
	go inc(&i, &mutex)
	time.Sleep(3 * time.Second)
	fmt.Println(i)
}

func inc(n *int, mutex *sync.Mutex) {
	for i := 0; i < 10000000; i++ {
		mutex.Lock()
		*n++
		mutex.Unlock()
	}
}
