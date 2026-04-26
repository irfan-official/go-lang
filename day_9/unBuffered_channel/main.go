package main

import "fmt"

func Worker(id int, myChan chan<- string) {
	fmt.Println("work done id: ", id)
	myChan <- "Success"
}

// blocking code
// func main() {
// 	myChan := make(chan string)

// 	go Worker(1, myChan)

// 	fmt.Print("start server ... \n")

// 	fmt.Println(<- myChan)

// 	fmt.Print("server closed gracefully ... \n")
// }



// unblocking code with select
func main(){
	myChan := make(chan string)
	defer close(myChan)

	go Worker(1, myChan)

	fmt.Print("start server ... \n")

	select {
	  case msg := <- myChan:
		fmt.Println(msg)
	}

	fmt.Print("server closed gracefully ... \n")

}