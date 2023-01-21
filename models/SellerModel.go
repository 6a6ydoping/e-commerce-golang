package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateSeller(firstName, lastName, email, password string) Seller {
	return Seller{FirstName: firstName, LastName: lastName, Email: email, Password: password}
}
