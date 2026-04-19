package routes 

import (
	"github.com/gin-gonic/gin"
	"my-app/modules/user"
)

func ServeRoutes(r *gin.Engine){
	routes := r.Group("/api/v1")

	routes.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	},
)

	user.UserRoutes(routes)
}