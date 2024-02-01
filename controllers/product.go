package controllers

import (
	"net/http"

	"github.com/JcsnP/postgres-docker/models"
	"github.com/gin-gonic/gin"
)

// get all products
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": products,
	})
}

// create product
func CreateProduct(c *gin.Context) {
	var input models.CreateProduct
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error", 
			"data": err.Error(),
		})
		return
	}

	// then create a product
	product := models.Product{ Title: input.Title, Description: input.Description }
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": product,
	})
}

// find a product
func FindProduct(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": product,
	})
}

// update product
func UpdateProduct(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var input models.UpdateProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": product,
	})
}

// delete product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}