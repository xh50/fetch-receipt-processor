package routes

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func ReceiptRoutes(router *gin.Engine) {
	router.GET("/", indexPage)
	router.GET("index", indexPage)
}

func indexPage(c *gin.Context) {
	log.Println("Call Index Page")
	c.String(http.StatusOK, "Index Page")
}