package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(work string, t int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(t))
	ch <- "work done " + work
	// fmt.Println("work done", work)
}

func main(){
	
	start := time.Now()

	var wg sync.WaitGroup

	channel_1 := make(chan string, 3)

	wg.Add(3)

	go Worker("img process 1", 10, &wg, channel_1)
	go Worker("img process 2", 100, &wg, channel_1)
	go Worker("img process 3", 30, &wg, channel_1)

	wg.Wait()
	close(channel_1)

	end := time.Since(start)


	// for i := 0; i < 3; i++{
	// 	result := <- channel_1
	// 	fmt.Println(result)
	// }

	for i := 0; i < 3; i++{
		fmt.Println(<- channel_1)
	}

	fmt.Println("Time taken: ", end)

}