package models

type RawItem struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Item struct {
	ShortDescription string
	Price float64
}