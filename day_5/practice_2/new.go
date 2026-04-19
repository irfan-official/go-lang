package practice2

import (
	"fmt"
	"math/rand"
)

// 1. init function

func init(){
	fmt.Println("Server starred")
	defer fmt.Println(rand.Int())
}

