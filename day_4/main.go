package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-app/routes"
)


func main(){
	fmt.Print("Hello world")
	server := gin.Default()

	routes.ServeRoutes(server)

	server.Run(":3000")
}