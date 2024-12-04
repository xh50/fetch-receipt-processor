package routes

import (
	"fetch-receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

func ReceiptRoutes(router *gin.Engine) {
	router.GET("/receipts/:id/points", controllers.GetReceiptByID)
	router.GET("/index/", controllers.IndexPage)
	router.GET("/", controllers.IndexPage)
}
