package main

import (
	"fmt"
	. "fmt"
)

// /
//
// simple function
func Hello() {
	fmt.Printf("hello")
}

type Myfucntion func(param1, param2 int) bool

// function expression
var NewFunction Myfucntion = func(p, q int) bool {
	fmt.Printf("value ==> %v %d", p, q)
	return !true
}

func init() {
	defer fmt.Println("Server closed gracefully")
	var test = NewFunction(1, 2)
	fmt.Println("\n", test)
}

// iife
// func int(){
// 	func (){
// 		fmt.Println("Hello its a iife like function")
// 	}()
// }


type Afunc func() int

var Count = func() Afunc {
	num := 0
	return func() int {
		num++
		return num
	}
}



func init() {
	
	cc := &Count

	test := (*cc)()

	Println((*&test)())
	Println((*&test)())
	Println((*&test)())
	Println((*&test)())


	// Println((*cc)()(), " ", (*cc)())
	// Println((*cc)()(), " ", (*cc)())
	// Println((*cc)()(), " ", (*cc)())
	// Println((*cc)()(), " ", (*cc)())

}

// type Afucntion func ()

// var HHello Afucntion = func(){
// 	Println("Hello")
// }

// func init(){
// 	v := &HHello
// 	(*v)()
// }