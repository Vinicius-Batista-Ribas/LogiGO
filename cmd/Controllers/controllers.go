package controllers

import (
	"LogiGO/cmd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {

	var orders []models.Pedido
	database.db.Find(&orders)
	c.JSON(http.StatusOK, orders)

}

func PostOrder(c *gin.Context) {
	var order models.Pedido
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.db.Create(&order)
	c.JSON(http.StatusCreated, order)
}
