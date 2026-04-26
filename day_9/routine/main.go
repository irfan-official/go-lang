package main

import (
	"fmt"
	"sync"
)

func Hello(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("-----> ", i)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("start")
	defer fmt.Println("server close gracefully")

	for i := 0; i < 3; i++ {
		wg.Add(i)

		go Hello(i, &wg)
	}
	fmt.Println("server close gracefully 1")
	wg.Wait()
	fmt.Println("server close gracefully 2")
}
