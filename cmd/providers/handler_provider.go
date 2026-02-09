package provider

import (
	controllers "LogiGO/cmd/controllers"
	handler "LogiGO/cmd/handler"
)

func NewPaymentHandlerMock(controller controllers.PaymentController) *handler.PaymentHandler {
	return handler.NewPaymentHandler(controller)
}
