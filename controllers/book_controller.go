package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iurylemos/webapi-golang/database"
	"github.com/iurylemos/webapi-golang/models"
)

func ShowBook(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"value": "ok",
	// })

	id := c.Param("id")

	// transform string to int
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}
	db := database.GetDataBase()

	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book" + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}