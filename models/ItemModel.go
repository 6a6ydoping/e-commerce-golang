package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	SellerID uint
	Name     string
	Price    float32
	Quantity uint32
	Rating   float32
}

func CreateItem(sellerID uint, itemName string, price float32, quantity uint32) Item {
	return Item{SellerID: sellerID, Name: itemName, Price: price, Quantity: quantity, Rating: 0}
}
