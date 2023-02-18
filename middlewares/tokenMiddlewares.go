package middlewares

import (
	"e-commerce-app/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

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
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})
}

func DeleteTokenFromCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return
		}
		fmt.Println(err)
		log.Fatal(err)
	}
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, cookie)
	fmt.Println("Cookie deleted")
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

func CreateClientToken(client *models.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   client.ID,
		"role": "Client",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
}

func CreateSellerToken(seller *models.Seller) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   seller.ID,
		"role": "Seller",
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWTSecretKey")))
}
