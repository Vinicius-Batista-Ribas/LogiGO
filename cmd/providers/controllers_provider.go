package provider

import (
	dto "LogiGO/cmd/dto"
	"LogiGO/cmd/models"
)

type ProviderMock struct {
	GetAllFn      func() ([]models.Payment, error)
	GetByTypeFn   func(models.PaymentType) ([]models.Payment, error)
	GetByMethodFn func(models.PaymentMethod) ([]models.Payment, error)
	CreateFn      func(*models.Payment) error
	GetSummaryFn  func() (*dto.PaymentSummaryDTO, error)
}

func (m *ProviderMock) GetAll() ([]models.Payment, error) {
	return m.GetAllFn()
}

func (m *ProviderMock) GetByType(t models.PaymentType) ([]models.Payment, error) {
	return m.GetByTypeFn(t)
}

func (m *ProviderMock) GetByMethod(method models.PaymentMethod) ([]models.Payment, error) {
	return m.GetByMethodFn(method)
}

func (m *ProviderMock) Create(p *models.Payment) error {
	return m.CreateFn(p)
}

func (m *ProviderMock) GetSummary() (*dto.PaymentSummaryDTO, error) {
	return m.GetSummaryFn()
}
