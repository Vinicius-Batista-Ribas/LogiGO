package handler

import (
	controllers "LogiGO/cmd/Controllers"
	dto "LogiGO/cmd/dto"
	"LogiGO/cmd/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	controller controllers.PaymentController
}

func NewPaymentHandler(controller controllers.PaymentController) *PaymentHandler {
	return &PaymentHandler{controller: controller}
}

func (h *PaymentHandler) GetPayments(c *gin.Context) {
	payments, err := h.controller.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetPaymentsByType(c *gin.Context) {
	paymentType := models.PaymentType(
		strings.ToUpper(c.Param("type")),
	)

	if !paymentType.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payment type",
		})
		return
	}

	payments, err := h.controller.GetByType(paymentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}
func (h *PaymentHandler) GetPaymentsByMethod(c *gin.Context) {
	method := models.PaymentMethod(
		strings.ToUpper(c.Param("method")),
	)

	if !method.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payment method",
		})
		return
	}

	payments, err := h.controller.GetByMethod(method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var dto dto.CreatePaymentDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := dto.ToModel()

	if err := h.controller.Create(payment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetSummary(c *gin.Context) {
	summary, err := h.controller.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
