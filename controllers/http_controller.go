package controllers

import (
	"fetch-receipt-processor/models"

	"github.com/gin-gonic/gin"
	
	"log"
	"net/http"
)

var holder = make(map[string]models.Receipt)

func IndexPage(c *gin.Context) {
	log.Println("Call Index Page")
	c.String(http.StatusOK, "Index Page")
}

func GetReceiptByID(c *gin.Context) {
	log.Println("Call GetReceiptByID")
	id := c.Param("id")
	log.Printf("Get ID: %s\n", id)
	c.String(http.StatusOK, id)
}
