package main

import "fmt"

var name string = "shadow"

func main() {
	// test()

	mao := "kao"

	val := *(&mao)

	fmt.Printf(val)

	// maker()

}