package controllers

import (
	"strconv"

	"github.com/YagoAyala/go-rest-api.git/database"
	"github.com/YagoAyala/go-rest-api.git/models"
	"github.com/YagoAyala/go-rest-api.git/services"
	"github.com/gin-gonic/gin"
)

func FindAllUser(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.User
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func FindOneUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.User
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find the user by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.Status(201)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditUser(c *gin.Context) {
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
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
