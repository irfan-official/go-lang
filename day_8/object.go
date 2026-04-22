package main

import (
	"fmt"
	_ "sync"
)

func init() {
	fmt.Println("this is the object")
}

type User struct {
	Name  string
	Email string
}

type Student struct {
	User
	ID int
}

type Teacher struct {
	User
	Subject string
}

func init() {
	autoID := Count()

	// Student1 := Student{
	// 	User: User{
	// 		Name:  "Irfan",
	// 		Email: "master@gmail.com",
	// 	},
	// 	ID: autoID(),
	// }

	// Student2 := Student{
	// 	User: User{
	// 		Name:  "Irfan",
	// 		Email: "master@gmail.com",
	// 	},
	// 	ID: autoID(),
	// }

	list := []Student{}

	for _, k := range "123" {
		list = append(list, Student{
			User: User{
				Name:  "Irfan",
				Email: "master@gmail.com",
			},
			ID: autoID(),
		},
		)
		fmt.Println(k)
	}
	fmt.Println(list)
}
