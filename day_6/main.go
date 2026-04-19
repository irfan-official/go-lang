package main

import (
	"fmt"
)

func main() {
	var res int = first_class(10, 10, func(a, b int) int {
		return a + b
	})

	fmt.Println(res)

	res2 := second_class(10, 10, func(a, b int) Return_type {
		return func (x, y int) int {
			return a + b + x + y
		}
	})

	judge1 := res2(10, 10)
	judge2 := res2(20, 10)
	judge3 := res2(30, 10)


	fmt.Println(judge1)
	fmt.Println(judge2)
	fmt.Println(judge3)
}
