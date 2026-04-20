package main

import (
	"fmt"
)

func intit() {
	fmt.Println("Server started")
	defer fmt.Println("Server close gracefully")
}

type Student struct {
	Name string
	Role int
}

var st1 Student = Student{
	Name: "Irfan",
	Role: 12,
}

// method
func (st1 Student) multiplereturn(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("error occur")
	}
	return num1 / num2, nil
}

// named function, naked return
func genName(suggestions string) (name string, age int) {
	name = suggestions
	age = 12
	return
}

type TakeMul struct {
	count     int
	total_val int
}

// variadic function
func takeMul(rest ...int) (total TakeMul) {

	total.count = 0
	total.total_val = 0

	for _, val := range rest {
		total.count += 1
		total.total_val += val
	}

	return
}

// Anonymous Functions (Function Literals)
// IIFE
func supportAnnoy() {
	func() {
		fmt.Println("I am anonymous")
	}()

	// anoy1 := func() {
	// 	fmt.Println("function2 has called")
	// }

	// anoy1()

	var anoy func() string
	
	anoy = func() string {
		return "mao mao"
	}

	var res3 string = anoy()

	fmt.Println("res3 => ", res3)
} 

// first class function
func first_class(a, b int, op func(int, int) int) int {
	return op(a, b)
}

type Return_type func(a, b int) int

func second_class(a, b int, op func (int, int) Return_type) Return_type {
	return op(a, b)
}