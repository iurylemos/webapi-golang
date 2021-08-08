package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/database"
	"github.com/iurylemos/webapi-golang/models"
	"github.com/iurylemos/webapi-golang/services"
)

func Login(c *gin.Context) {
	db := database.GetDataBase()

	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User

	dbError := db.Where("email = ?", login.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user: " + dbError.Error(),
		})
		return
	}

	// check if password is equal password of user database
	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
