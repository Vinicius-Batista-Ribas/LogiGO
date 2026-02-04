package repository

// Repository Used to conected with the database
import (
	"LogiGO/cmd/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	FindAll() ([]models.Payment, error)
	GetPaymentsByMethod(method string) ([]models.Payment, error)
	GetPaymentsByType(type_ string) ([]models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) FindAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetPaymentsByMethod(method string) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("method = ?", method).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetPaymentsByType(type_ string) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("type = ?", type_).Find(&payments).Error
	return payments, err
}
