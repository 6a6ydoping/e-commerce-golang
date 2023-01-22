package middlewares

import (
	"e-commerce-app/helpers"
	"e-commerce-app/initializers"
	"e-commerce-app/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
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

func CreateSellerToken(seller *models.Seller) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"seller_id": seller.ID,
		"role":      "Seller",
		"exp":       time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
}

func CreateClientToken(client *models.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"seller_id": client.ID,
		"role":      "Client",
		"exp":       time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
}

func GetSellerByEmail(email string) (*models.Seller, error) {
	var seller models.Seller
	err := initializers.DB.Model(models.Seller{Email: email}).First(&seller).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.UserNotFoundError{Message: "User not found"}
		}
		log.Fatal("Get seller by email crashed")
	}
	return &seller, nil
}

func MakeAuthRequest(token string) error {
	req, err := http.NewRequest("POST", "http://localhost:8000/private", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	httpClient := http.Client{}
	_, err = httpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
