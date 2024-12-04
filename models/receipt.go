package models

import (
	"time"
)

type RawReceipt struct {
	ID string
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []RawItem `json:"items"`
	Total string `json:"total"` 
}

type Receipt struct {
	ID string
	Retailer string 
	PurchaseDate time.Time
	PurchaseTime time.Time
	Items []Item 
	Total float64
}