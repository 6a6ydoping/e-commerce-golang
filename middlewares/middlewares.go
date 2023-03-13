package middlewares

import (
	"e-commerce-app/helpers"
	"e-commerce-app/initializers"
	"e-commerce-app/models"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckEmailInDB(email, userType string) error { //TODO: разбить на две функции мб
	if userType == "Client" {
		var client models.Client
		err := initializers.DB.Where("email = ?", email).First(&client).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
		}
		return fmt.Errorf("Email already exists")
	}

	if userType == "Seller" {
		var seller models.Seller
		err := initializers.DB.Where("email = ?", email).First(&seller).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
		}
		return fmt.Errorf("Email already exists")
	}

	return nil
}

func InsertSellerIntoDataBase(s models.Seller) error {
	result := initializers.DB.Create(&s)
	return result.Error
}

func InsertClientIntoDataBase(c models.Client) error {
	result := initializers.DB.Create(&c)
	return result.Error
}

func InsertItemIntoDataBase(i models.Item) error {
	result := initializers.DB.Create(&i)
	return result.Error
}

func CheckEmailAndPasswordInDB(email, password, userType string) error {
	if userType == "Seller" {
		var seller = models.Seller{}
		err := initializers.DB.Where("email = ?", email).First(&seller).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &helpers.UserNotFoundError{Message: "User not found"}
			}
		}
		return bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(password))

	} else if userType == "Client" {
		var client = models.Client{}
		err := initializers.DB.Where("email = ?", email).First(&client).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &helpers.UserNotFoundError{Message: "User not found"}
			}
		}
		return bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	}
	return nil
}

func GetSellerByEmail(email string) (*models.Seller, error) {
	var seller models.Seller
	err := initializers.DB.Where("email = ?", email).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.UserNotFoundError{Message: "User not found"}
		}
		log.Fatal("Get seller by email crashed")
	}
	return &seller, nil
}

func GetClientByEmail(email string) (*models.Client, error) {
	var client models.Client
	err := initializers.DB.Where("email = ?", email).First(&client).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.UserNotFoundError{Message: "User not found"}
		}
		log.Fatal("Get client by email crashed")
	}
	return &client, nil

}

func GetSellerByID(id interface{}) (*models.Seller, error) {
	var seller models.Seller
	err := initializers.DB.Model(models.Seller{}).First(&seller, fmt.Sprint(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.UserNotFoundError{Message: "User not found"}
		}
		fmt.Println(err)
		log.Fatal("Get seller by id crashed")
	}
	return &seller, nil
}

func GetClientByID(id interface{}) (*models.Client, error) {
	var client models.Client
	err := initializers.DB.Model(models.Client{}).First(&client, fmt.Sprint(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.UserNotFoundError{Message: "User not found"}
		}
		fmt.Println(err)
		log.Fatal("Get client by id crashed")
	}
	return &client, nil
}

func GetSellingItems(sellingItems *[]models.Item, itemName string) error {
	if itemName != "" {
		if err := initializers.DB.Where("name LIKE ?", itemName+"%").Find(&sellingItems).Error; err != nil {
			return errors.New("failed to fetch all selling items")
		}
	} else if err := initializers.DB.Find(&sellingItems).Error; err != nil {
		return errors.New("failed to fetch all selling items")
	}
	return nil
}
