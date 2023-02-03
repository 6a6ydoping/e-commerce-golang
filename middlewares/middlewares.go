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

func CreateSellerToken(seller *models.Seller) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   seller.ID,
		"role": "Seller",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
}

func CreateClientToken(client *models.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   client.ID,
		"role": "Client",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
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

func GetRoleFromStringToken(tokenString string) (interface{}, error) {
	token, err := decodeStringToken(tokenString)
	if err != nil {
		fmt.Println("error while decoding token string")
		log.Fatal(err)
		return "", err
	}
	//Вытягиваем инфу юзера из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("error while getting info from users token")
		log.Fatal(err)
	}
	return claims["role"], nil
}

func GetIDFromStringToken(tokenString string) (interface{}, error) {
	token, err := decodeStringToken(tokenString)
	if err != nil {
		fmt.Println("error while decoding token string")
		log.Fatal(err)
		return "", err
	}
	//Вытягиваем инфу юзера из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("error while getting info from users token")
		log.Fatal(err)
	}
	return claims["id"], nil
}

func decodeStringToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTSecretKey")), nil
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return token, nil
}

func GetTokenValueFromCookie(r *http.Request) (string, error) {
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", err
		}
	}
	return token.Value, err
}

func SetTokenToCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
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

func CheckTokenExpiry(tokenString string) (bool, error) {
	// Поменять на DecodeStringToken()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTSecretKey")), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return false, nil
		}
		return true, nil
	}
	return false, fmt.Errorf("invalid token")
}
