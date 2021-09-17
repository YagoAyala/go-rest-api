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
			user.GET("/:id", controllers.FindOneUser)
			user.POST("/", controllers.CreateUser)
			user.PUT("/", controllers.EditUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.FindAllBook)
			books.GET("/:id", controllers.FindOneBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		main.POST("login", controllers.Login)
	}
	//TODO: arrumar o erro em relação com a autenticação do middleware
	return router
}
