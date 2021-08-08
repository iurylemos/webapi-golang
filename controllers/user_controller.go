package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/database"
	"github.com/iurylemos/webapi-golang/models"
	"github.com/iurylemos/webapi-golang/services"
)

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// encoding 256 password
	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
