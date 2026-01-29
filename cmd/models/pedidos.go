package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Pedido struct {
	gorm.Model

	ID            uint `gorm:"primaryKey;autoIncrement"`
	customerName  string
	customerPhone string

	pickupAddress   string
	deliveryAddress string

	status     string
	createedAt time.Time
	updatedAt  time.Time
}

type OrderStatus string

const (
	OrderCreated   OrderStatus = "CREATED"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderInTransit OrderStatus = "IN_TRANSIT"
	OrderDelivered OrderStatus = "DELIVERED"
	OrderCanceled  OrderStatus = "CANCELED"
)

func ValidateOrder(order *Pedido) error {
	if err := validator.Validate(order); err != nil {
		return err
	}
	return nil
}
