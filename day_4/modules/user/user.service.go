package user

import (
	"fmt"
)

func ServiceGetUsers() string {
	fmt.Println("service called !")
	return "completed"
}
