package main

import "fmt"

// 1. Basic functions
func Add(a, b int) int{
	return a+b
}

// 2. function with mutiple return 
func Divide(a, b int) (int, error){
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil
}

// 3. named return functions
func Rectangle(w, h int) (area int) {
	area = w * h
	return
} 

// 4. variadic functions
func Sum (rest ...int) (length int) {
	length = len(rest)
	return
}

 // 5. Anonymous functions
 func inti(){
	func(){
		fmt.Println("Initializing...")
	}()
 }

