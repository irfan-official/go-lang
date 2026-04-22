package main

import "fmt"

func main() {
	func(p, q int) {
		fmt.Println("Hello")
	}(10, 20)
}
