package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateClient(firstName, lastName, email, password string) Client {
	return Client{FirstName: firstName, LastName: lastName, Email: email, Password: password}
}
