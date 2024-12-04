package services

import (
	"fetch-receipt-processor/models"

	"log"
	"unicode"
	"math"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
	log.Printf("Calculate point for receipt ID: %s\n", receipt.ID)
	point := 0
	point += retailNamePoint(receipt)
	point += roundDollarAmountPoint(receipt)
	point += multiple25Point(receipt)
	point += twoItemsPoint(receipt)
	point += trimmedItemDescription(receipt)
	point += oddDatePoint(receipt)
	point += between2and4(receipt)
	return point
}

/*
	One point for every alphanumeric character in the retailer name.
*/
func retailNamePoint(r models.Receipt) int {
	earnedPoint := 0
	for _, r := range r.Retailer {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			earnedPoint++
		}
	}
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}

/*
	50 points if the total is a round dollar amount with no cents.
*/
func roundDollarAmountPoint(r models.Receipt) int {
	earnedPoint := 0
	if math.Mod(r.Total, 1) == 0 {
		earnedPoint = 50
	}
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}

/*
	25 points if the total is a multiple of 0.25.
*/
func multiple25Point(r models.Receipt) int {
	earnedPoint := 0
	if math.Mod(math.Mod(r.Total, 0.25), 1) == 0 {
		earnedPoint = 25
	} 
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}

/*
	5 points for every two items on the receipt.
*/
func twoItemsPoint(r models.Receipt) int {
	earnedPoint := 5 *(len(r.Items) / 2)
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}

/*
	If the trimmed length of the item description is a multiple of 3, 
	multiply the price by 0.2 and round up to the nearest integer. 
	The result is the number of points earned.
*/
func trimmedItemDescription(r models.Receipt) int {
	earnedPoint := 0
	for _, item := range r.Items {
		trimmedItemDescription := strings.TrimRight(strings.TrimLeft(item.ShortDescription, " "), " ")
		if len(trimmedItemDescription) % 3 == 0 {
			temp := 0.2 * item.Price
			round := int(temp) + 1
			earnedPoint += round
		} 
	}
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint	
}

/*
	6 points if the day in the purchase date is odd.
*/
func oddDatePoint(r models.Receipt) int {
	earnedPoint := 0
	_, _, day := r.PurchaseDate.Date()
	if day % 2 != 0 {
		earnedPoint = 6
	}
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}

/*
	10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/
func between2and4(r models.Receipt) int {
	timeStr := "15:04"
	earlyLimit := "14:00"
	lateLimit := "16:00"
	earlyTime, err1 := time.Parse(timeStr, earlyLimit)
	lateTime, err2 := time.Parse(timeStr, lateLimit)
	earnedPoint := 0
	if err1 != nil || err2 != nil {
		log.Println("Parse Time Error:", err1, err2)
		return 0
	}
	if r.PurchaseTime.After(earlyTime) && r.PurchaseTime.Before(lateTime) {
		earnedPoint = 10
	}
	log.Printf("EarnedPoint: %d\n", earnedPoint)
	return earnedPoint
}