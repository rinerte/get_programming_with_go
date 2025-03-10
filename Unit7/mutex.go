package main

import (
	"fmt"
	"sync"
)

func main() {
	i := 0
	var mutex sync.Mutex

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		inc(&i, &mutex)
		wg.Done()
	}()
	go func() {
		inc(&i, &mutex)
		wg.Done()
	}()
	go func() {
		inc(&i, &mutex)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(i)
}

func inc(n *int, mutex *sync.Mutex) {
	for i := 0; i < 100000000; i++ {
		mutex.Lock()
		*n++
		mutex.Unlock()
	}
}
