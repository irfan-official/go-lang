package main

import (
	// "fmt"
	"math/rand"
	"github.com/gin-gonic/gin"
	"github.com/irfan-official/go-lang/test"
)

func welcomeUser(c *gin.Context){
	c.JSON(200, gin.H{
		"success": true,
		"num": rand.Int(),
		"message": test.Message,
	})
}

func main(){
	route := gin.Default()
	route.GET("/", welcomeUser)
	route.Run(":3000")
}