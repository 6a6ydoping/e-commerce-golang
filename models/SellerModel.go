package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateSeller(firstName, lastName, email, password string) Seller {
	return Seller{FirstName: firstName, LastName: lastName, Email: email, Password: password}
}
