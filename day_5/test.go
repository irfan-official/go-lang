package main

import (
	"fmt"
)

type Student struct {
	ID    int
	Name  string
	Email string
}

var user Student = Student{
	ID:    1,
	Name:  "Irfan",
	Email: "irfan@gmail.com",
}

func test() {

	user.Email = "mao"

	fmt.Println(user)
}

func (user Student) divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	
	result := operate(2, 3, func(x, y int) int {
	return x + y
})

fmt.Println(result)

	test()
	res, err := user.divide(10, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

type newFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

var f newFunc = add

// start immidiately before start the main
func init() {
	defer fmt.Println("server closed gracefully")
	fmt.Println("Server started")
	fmt.Println(add(10, 2))
}

