package handler_test

import (
	dto "LogiGO/cmd/dto"
	handler "LogiGO/cmd/handler"
	"LogiGO/cmd/models"
	controllersMock "LogiGO/cmd/providers"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewTestPaymentHandler(mock *controllersMock.ProviderMock) *handler.PaymentHandler {
	return handler.NewPaymentHandler(mock)
}

func TestGetPayments_Sucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		GetAllFn: func() ([]models.Payment, error) {
			return []models.Payment{
				{Amount: 100},
				{Amount: 200},
			}, nil
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments", handler.GetPayments)

	req, _ := http.NewRequest(http.MethodGet, "/payments", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetPayments_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		GetAllFn: func() ([]models.Payment, error) {
			return nil, errors.New("db error")
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments", handler.GetPayments)

	req, _ := http.NewRequest(http.MethodGet, "/payments", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestGetPaymentsByType_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockController := &controllersMock.ProviderMock{
		GetByTypeFn: func(paymentType models.PaymentType) ([]models.Payment, error) {
			assert.Equal(t, models.PaymentType("EXPENSE"), paymentType)
			return []models.Payment{
				{Amount: 100},
			}, nil
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/type/:type", handler.GetPaymentsByType)

	req, _ := http.NewRequest(http.MethodGet, "/payments/type/EXPENSE", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetPaymentsByType_InvalidType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/type/:type", handler.GetPaymentsByType)

	req, _ := http.NewRequest(http.MethodGet, "/payments/type/invalid", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestGetPaymentsByMethod_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockController := &controllersMock.ProviderMock{
		GetByMethodFn: func(method models.PaymentMethod) ([]models.Payment, error) {
			assert.Equal(t, models.PaymentMethod("CREDIT"), method)
			return []models.Payment{
				{Amount: 150},
			}, nil
		},
	}
	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/method/:method", handler.GetPaymentsByMethod)

	req, _ := http.NewRequest(http.MethodGet, "/payments/method/CREDIT", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetPaymentsByMethod_InvalidMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/method/:method", handler.GetPaymentsByMethod)

	req, _ := http.NewRequest(http.MethodGet, "/payments/method/invalid", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreatePayment_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		CreateFn: func(p *models.Payment) error {
			// valida se o handler montou o model corretamente
			assert.Equal(t, 100.0, p.Amount)
			assert.Equal(t, models.Credit, p.Method)
			assert.Equal(t, models.Expense, p.Type)
			assert.Equal(t, "Almoço", p.Description)
			return nil
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.POST("/payments", handler.CreatePayment)

	body := `{
		"date": "2026-02-01T10:00:00Z",
		"amount": 100,
		"description": "Almoço",
		"method": "CREDIT",
		"type": "EXPENSE"
	}`

	req, _ := http.NewRequest(http.MethodPost, "/payments", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestCreatePayment_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewTestPaymentHandler(&controllersMock.ProviderMock{})

	router := gin.Default()
	router.POST("/payments", handler.CreatePayment)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/payments",
		strings.NewReader(`{ invalid json`),
	)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreatePayment_ControllerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		CreateFn: func(p *models.Payment) error {
			return errors.New("validation error")
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.POST("/payments", handler.CreatePayment)

	body := `{
		"date": "2026-02-01T10:00:00Z",
		"amount": 100,
		"description": "Almoço",
		"method": "CREDIT",
		"type": "EXPENSE"
	}`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/payments",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestGetSummary_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		GetSummaryFn: func() (*dto.PaymentSummaryDTO, error) {
			return &dto.PaymentSummaryDTO{
				Income:  1000,
				Expense: 400,
				Balance: 600,
			}, nil
		},
	}
	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/summary", handler.GetSummary)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/payments/summary",
		nil,
	)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.JSONEq(t, `
	{
		"income": 1000,
		"expense": 400,
		"balance": 600
	}
	`, resp.Body.String())
}

func TestGetSummary_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockController := &controllersMock.ProviderMock{
		GetSummaryFn: func() (*dto.PaymentSummaryDTO, error) {
			return nil, errors.New("db error")
		},
	}

	handler := NewTestPaymentHandler(mockController)

	router := gin.Default()
	router.GET("/payments/summary", handler.GetSummary)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/payments/summary",
		nil,
	)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
