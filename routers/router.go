package routers

import (
	"go-jwt/controller"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegistration)
		userRouter.POST("/login", controller.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authorization())
		productRouter.POST("/", controller.CreateProduct)
	}
	return r
}
