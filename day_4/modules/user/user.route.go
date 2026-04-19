
package user


import ("github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup){

	route := rg.Group("/users")

	route.GET("/", GetUsers)

}