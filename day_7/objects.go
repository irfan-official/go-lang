package main

import "fmt"

// import ("fmt")

type Student struct{
	Name string
	Id   int
}

var s1 Student = Student{
	Name: "Irfan",
}


func (s1 Student) Holder(num int) int{
	return num
}

func test(){
	fmt.Println(s1)
	s1.Holder(10)
}

// Anonymous Structs

func maker(){
	user := struct{
		Name string
	}{Name: "Irfan"}

	fmt.Printf(user.Name)
}

// struct embedding

type User struct {
	Name string
	ID int
}

type Admin struct{
	User
	Level int
}

func test_embedding(){
	a := Admin{
		User: User{Name: "Alice", ID: 1},
		Level: 10,
	}

	fmt.Println(a)
}