package main

import ("fmt"
"github.com/gin-gonic/gin"
 "backdnd-template/routes"
)


func main(){
	fmt.Print("✅ Server started")

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run(":3000")

}