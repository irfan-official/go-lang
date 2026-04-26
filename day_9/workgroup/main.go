package main

import (
	"fmt"
	"sync"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("finished task ... %v\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go Worker(1, &wg)
	go Worker(2, &wg)

	fmt.Println("main task completed")

	wg.Wait()
}
