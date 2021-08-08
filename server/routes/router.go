package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
		}
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}
	}

	return router
}
