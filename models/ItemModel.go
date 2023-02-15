package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	SellerID uint    `json:"sellerID"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint32  `json:"quantity"`
	Rating   float32 `json:"rating"`
}

func CreateItem(sellerID uint, itemName string, price float32, quantity uint32) Item {
	return Item{SellerID: sellerID, Name: itemName, Price: price, Quantity: quantity, Rating: 0}
}	
