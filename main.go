package main

import (
    "fmt"
	"math/rand"

    // local
    "github.com/irfan-official/go-lang/test"

    // external
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println(test.Tog(10, 20))

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"success": true, "num": rand.Int(), "message": test.Message})
    })
    r.Run()
}