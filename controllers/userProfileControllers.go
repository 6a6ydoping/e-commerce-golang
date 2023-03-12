package controllers

import (
	"e-commerce-app/helpers"
	"e-commerce-app/middlewares"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetUserProfileInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		helpers.EnableCors(&w)
		fmt.Println(r.Cookie("token"))
		token, err := middlewares.GetTokenValueFromCookie(r)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		userID, err := middlewares.GetIDFromStringToken(token)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		role, err := middlewares.GetRoleFromStringToken(token)
		if role == "Seller" && err == nil {
			fmt.Println("OK")
			seller, err := middlewares.GetSellerByID(userID)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			jsonData, err := json.Marshal(seller)
			if err != nil {
				log.Fatal("couldnt json data")
				return
			}
			_, err = w.Write(jsonData)
			if err != nil {
				fmt.Println("error in getUserProf controller")
			}
		} else if role == "Client" && err == nil {
			client, err := middlewares.GetClientByID(userID)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			jsonData, err := json.Marshal(client)
			if err != nil {
				log.Fatal("couldnt json data")
				return
			}
			_, err = w.Write(jsonData)
			if err != nil {
				fmt.Println("error in getUserProf controller")
			}
		} else {
			fmt.Println(err)
		}
	}

}
