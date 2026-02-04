package routes

import (
	controllers "LogiGO/cmd/Controllers"
	handler "LogiGO/cmd/Handler"
	"LogiGO/cmd/database"
	"LogiGO/cmd/repository"

	"github.com/gin-gonic/gin"
)

func HandlerRequest() {
	repo := repository.NewPaymentRepository(database.DB)
	controller := controllers.NewPaymentController(repo)
	handler := handler.NewPaymentHandler(controller)

	r := gin.Default()

	rotas := r.Group(payment)
	{
		rotas.GET("", handler.GetPayments)
		rotas.GET(paymentsByMethod, handler.GetPaymentsByMethod)
		rotas.GET(paymentsByType, handler.GetPaymentsByType)
		rotas.GET(summary, handler.GetSummary)

		rotas.POST("", handler.CreatePayment)
	}

	r.Run(":8080")
}
