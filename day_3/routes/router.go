package routes

import (
"github.com/gin-gonic/gin"
"backend-template/modules/user"
"backend-template/modules/waitlist"
)

func SetupRoutes(r *gin.Engine){

	api_V1 := r.Group("/api/v1")

	user.RegisterUserRoutes(api_V1)
	waitlist.RegisterWaitlistRoutes(api_V1)

}