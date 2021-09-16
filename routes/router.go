package routes

import (
	"github.com/YagoAyala/go-rest-api.git/controllers"
	"github.com/YagoAyala/go-rest-api.git/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.FindAll)
			books.GET("/:id", controllers.FindOne)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.Edit)
			books.DELETE("/:id", controllers.Delete)
		}

		main.POST("login", controllers.Login)
	}

	return router
}
