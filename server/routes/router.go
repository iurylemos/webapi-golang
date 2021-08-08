package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/controllers"
	"github.com/iurylemos/webapi-golang/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books", middlewares.Auth())
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

		main.POST("login", controllers.Login)
	}

	return router
}
