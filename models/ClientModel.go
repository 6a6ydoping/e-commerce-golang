package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateClient(firstName, lastName, email, password string) Client {
	return Client{FirstName: firstName, LastName: lastName, Email: email, Password: password}
}
