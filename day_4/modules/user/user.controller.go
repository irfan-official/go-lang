package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context){
	res := ServiceGetUsers()

	fmt.Println(res)

	c.JSON(200, gin.H{
		"success": true,
		"message": "hello",
	})
}