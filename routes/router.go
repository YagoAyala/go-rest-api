package routes

import (
	"github.com/YagoAyala/go-rest-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.FindAll)
			books.GET("/:id", controllers.FindOne)
			books.POST("/", controllers.Create)
			books.PUT("/", controllers.Edit)
			books.DELETE("/:id", controllers.Delete)
		}
	}

	return router
}
