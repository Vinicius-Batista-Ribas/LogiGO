package routes

import (
	Controllers "LogiGO/cmd/Controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequest() {

	r := gin.Default()

	r.GET("/order", Controllers.GetPayments)
	r.POST("/order", Controllers.PostPayment)
	r.Run()
}
