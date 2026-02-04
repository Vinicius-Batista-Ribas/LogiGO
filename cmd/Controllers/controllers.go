package controllers

import (
	"LogiGO/cmd/dto"
	err "LogiGO/cmd/erros"
	"LogiGO/cmd/models"
	"LogiGO/cmd/repository"
	"strings"
)

type PaymentController interface {
	Create(payment *models.Payment) error
	GetAll() ([]models.Payment, error)
	GetByMethod(pm models.PaymentMethod) ([]models.Payment, error)
	GetByType(pt models.PaymentType) ([]models.Payment, error)
	GetSummary() (*dto.PaymentSummaryDTO, error)
}

type paymentController struct {
	repo repository.PaymentRepository
}

func NewPaymentController(repo repository.PaymentRepository) PaymentController {
	return &paymentController{repo: repo}
}

func (c *paymentController) Create(payment *models.Payment) error {
	if payment.Amount <= 0 {
		return err.ErrInvalidAmount
	}
	return c.repo.Create(payment)
}

func (c *paymentController) GetAll() ([]models.Payment, error) {
	return c.repo.FindAll()
}

func (c *paymentController) GetByMethod(pm models.PaymentMethod) ([]models.Payment, error) {

	method := strings.ToUpper(string(pm))

	m := models.PaymentMethod(method)
	if !m.IsValid() {
		return nil, err.ErrInvalidStatus
	}

	return c.repo.GetPaymentsByMethod(method)
}

func (c *paymentController) GetByType(pt models.PaymentType) ([]models.Payment, error) {
	types := strings.ToUpper(string(pt))

	p := models.PaymentType(types)
	if !p.IsValid() {
		return nil, err.ErrInvalidType
	}
	return c.repo.GetPaymentsByType(types)
}

func (c *paymentController) GetSummary() (*dto.PaymentSummaryDTO, error) {
	payments, err := c.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var income, expense float64

	for _, p := range payments {
		switch p.Type {
		case models.Income:
			income += p.Amount
		case models.Expense:
			expense += p.Amount
		}
	}

	return &dto.PaymentSummaryDTO{
		Income:  income,
		Expense: expense,
		Balance: income - expense,
	}, nil
}
