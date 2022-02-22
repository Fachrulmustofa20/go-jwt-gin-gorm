package routers

import (
	"go-jwt/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegistration)
		userRouter.POST("/login", controller.UserLogin)
	}
	return r
}
