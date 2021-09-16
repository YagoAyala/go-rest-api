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
			user.GET("/", controllers.FindAllUser)
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.FindAllBook)
			books.GET("/:id", controllers.FindOne)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.Edit)
			books.DELETE("/:id", controllers.Delete)
		}

		main.POST("login", controllers.Login)
	}
	//TODO: arrumar o erro em relação com a autenticação do middleware
	return router
}
