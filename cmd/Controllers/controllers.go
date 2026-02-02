package controllers

import (
	"LogiGO/cmd/database"
	"LogiGO/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPayments(c *gin.Context) {

	var payments []models.Payment
	database.DB.Find(&payments)
	c.JSON(http.StatusOK, payments)
}

func PostPayment(c *gin.Context) {
	var order models.Payment
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidatePayment(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}
