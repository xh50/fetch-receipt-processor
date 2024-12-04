package controllers

import (
	"fetch-receipt-processor/models"
	"fetch-receipt-processor/utils"
	"fetch-receipt-processor/services"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
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
	if receipt, ok := holder[id]; ok {
		log.Printf("Found receipt for ID: %s\n", receipt.ID)
		point := services.CalculatePoints(receipt)
		temp := make(map[string]int)
		temp["point"] = point
		response, err := json.Marshal(temp)
		if err != nil {
			log.Println("Build JSON error: ", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, string(response))
	} else {
		log.Printf("Do not find receipt for ID: %s\n", id)
		c.JSON(http.StatusNotFound, "Do not find receipt with given ID")
	}
}

func GetAllReceipts(c *gin.Context) {
	c.JSON(http.StatusOK, holder)
}

func CreateReceipt(c *gin.Context) {
	var rawReceipt models.RawReceipt
	log.Println("Get receipt raw data. Start parsing..")
	if err := c.ShouldBindJSON(&rawReceipt); err != nil {
		log.Fatalln("Raw data parsing error.", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Raw data parsing error.": err.Error()})
		return
	}
	receipt, err := transferReceipt(rawReceipt)
	if err != nil {
		log.Fatalln("Receipt transfer error.", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Receipt transfer error.": err.Error()})
		return 
	}

	response, err := buildResponse(receipt)
	if err != nil {
		log.Fatalln("Build response error.", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Build response error.": err.Error()})
		return
	}
	log.Println("Create receipt successfully.")
	c.JSON(http.StatusOK, response)
}

func buildResponse(receipt models.Receipt) (string, error){
	jsonMap := map[string]interface{}{
		"ID": receipt.ID,
	}
	json, err := json.Marshal(jsonMap)
	return string(json), err
}

func transferReceipt(rawReceipt models.RawReceipt) (models.Receipt, error) {
	var receipt models.Receipt
	receipt.Retailer = rawReceipt.Retailer
	price, err := strconv.ParseFloat(rawReceipt.Total, 64)
	if err != nil {
		log.Fatalln("Price transfer error.", err.Error())
		return receipt, err
	}
	receipt.Total = price
	parsedDate, err := time.Parse(time.DateOnly, rawReceipt.PurchaseDate)
	if err != nil {
		log.Fatalln("Date parsing error.", err.Error())
		return receipt, err
	}
	receipt.PurchaseDate = parsedDate
	var timeStr = "15:04"
	parseTime, err := time.Parse(timeStr, rawReceipt.PurchaseTime)
	if err != nil {
		log.Fatalln("Time parsing error.", err.Error())
		return receipt, err
	}
	receipt.PurchaseTime = parseTime

	// transfer item data objects
	for i := range rawReceipt.Items {
		var item models.Item
		rawItem := &rawReceipt.Items[i]
		p, e := strconv.ParseFloat(rawItem.Price, 64)
		if e != nil {
			log.Println("Item transfer error.", e.Error())
			return receipt, e
		}
		item.Price = p
		item.ShortDescription = rawItem.ShortDescription
		receipt.Items = append(receipt.Items, item)
	}

	receipt.ID = utils.GenerateID().String()
	holder[receipt.ID] = receipt
	log.Println("Receipt data object transfer finished.")
	return receipt, nil
}