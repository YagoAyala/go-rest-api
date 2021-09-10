package controllers

import (
	"strconv"

	"github.com/YagoAyala/go-rest-api.git/database"
	"github.com/YagoAyala/go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Book
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func FindOne(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Book
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}
	// teste
	c.JSON(200, p)
}

func Create(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(201, p)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func Edit(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Book

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
