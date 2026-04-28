package main

import "fmt"

// 1. Basic functions
func Add(a, b int) int {
	return a + b
}

// 2. function with mutiple return
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 3. named return functions
func Rectangle(w, h int) (area int) {
	area = w * h
	return
}

// 4. variadic functions
func Sum(rest ...int) (length int) {
	length = len(rest)
	return
}

// 5. Anonymous functions
func inti() {
	func() {
		fmt.Println("Initializing...")
	}()
}

// 6. Functions as value
func init() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum of 3 and 5 is:", add(3, 5))
}

// 7. Higher-Order functions
func Operrate(a, b int, op func(int, int) int) int {
	return func (a, b int) int {
		return a * b
	}(a, b)
}

// 8. Clouser
func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 9. go routines function
func init() {
	go func() {
		fmt.Println("Hello from a goroutine!")
	}()
}

// 10. Methods
type Person struct {
	Name string
}

// Method with value receiver
func (u Person) SayHello() string {
	return "Hello" + u.Name
}

// Method with pointer value
func (u *Person) SetName(name string){
	u.Name = name
}