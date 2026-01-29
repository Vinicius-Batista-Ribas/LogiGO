package routes

import (
	Controllers "LogiGO/cmd/Controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequest() {

	r := gin.Default()

	r.GET("/order", Controllers.GetOrder)
	r.POST("/order", Controllers.PostOrder)
	r.Run()
}
