package main

import "fmt"

func main() {
	myChannel := make(chan string, 3)

	myChannel <- "Job 1 ..."
	myChannel <- "Job 2 ..."
	myChannel <- "Job 3 ..."

	close(
		myChannel,
	)

	fmt.Println(<- myChannel)
	fmt.Println(<- myChannel)
	fmt.Println(<- myChannel)
}