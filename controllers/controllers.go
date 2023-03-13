package controllers

import (
	"e-commerce-app/helpers"
	"e-commerce-app/middlewares"
	"e-commerce-app/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/home.html")
		if err != nil {
			log.Fatal("Error parsing file...")
		}
		t.Execute(w, nil)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	if r.Method == "GET" {
		// Берем токен
		token, err := middlewares.GetTokenValueFromCookie(r)
		if err != nil || token == "" {
			http.Error(w, "You have to login!", http.StatusUnauthorized)
			return
		}

		// Проверяем не истек ли срок годности токена
		isValid, err := middlewares.CheckTokenExpiry(token)
		if !isValid {
			http.Error(w, "Your session has ended, please login", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Smth bad happened", http.StatusBadGateway)
			return
		}

		//Получаем роль из токена
		role, err := middlewares.GetRoleFromStringToken(token)
		if err != nil {
			fmt.Println("Error get role from string token")
		}
		//Проверяем есть ли доступ к странице
		//TODO: Сделать мапу с ролями и их доступами
		if role == "Seller" || role == "Admin" {
			t, err := template.ParseFiles("views/createItem.html")
			if err != nil {
				log.Fatal("Error parsing file...")
			}
			t.Execute(w, nil)
		} else {
			http.Error(w, "You have no permission", http.StatusBadGateway)
		}

	} else if r.Method == "POST" {
		// Получаем токен селлера
		token, err := middlewares.GetTokenValueFromCookie(r)
		if err != nil {
			log.Fatal(err)
		}

		//Проверяем срок годности токена
		isValid, err := middlewares.CheckTokenExpiry(token)
		if !isValid {
			http.Error(w, "Your session has ended, please login", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Smth bad happened", http.StatusBadGateway)
			return
		}
		//Получаем айдишку из токена
		id, err := middlewares.GetIDFromStringToken(token)
		fmt.Println(id)

		//Ищем seller_id в бд
		seller, err := middlewares.GetSellerByID(id)
		if err != nil {
			http.Error(w, "Please login again", http.StatusBadGateway)
		}
		fmt.Println(seller.Email)
		// Парсим форму
		r.ParseForm()
		itemName := r.FormValue("itemName")
		stringPrice := r.FormValue("price")
		stringQuantity := r.FormValue("quantity")
		// Меняем типы (СДЕЛАТЬ НОРМАЛЬНУЮ ФУНКЦИЮ)
		price, err := strconv.ParseFloat(stringPrice, 32)
		if err != nil {
			fmt.Printf("Error type cast from str to f64")
		}
		var floatPrice float32 = float32(price)

		ui64, err := strconv.ParseUint(stringQuantity, 10, 64)
		if err != nil {
			panic(err)
		}

		uintQuantity := uint32(ui64)

		//Добавляем айтем в таблицу айтем с селлер айди
		err = middlewares.InsertItemIntoDataBase(models.CreateItem(seller.ID, itemName, floatPrice, uintQuantity))
		if err != nil {
			http.Error(w, "Failed to add item", http.StatusBadGateway)
			fmt.Println("Failed to add item")
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

}

func GetSellingItems(w http.ResponseWriter, r *http.Request) {
	queryItemName := r.URL.Query().Get("query")
	fmt.Println(queryItemName)
	//queryOrderBy := r.URL.Query().Get("orderBy")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		fmt.Println("Error here")
		return
	}
	defer r.Body.Close()
	fmt.Println(string(body))
	helpers.EnableCors(&w)
	var sellingItems []models.Item
	if err := middlewares.GetSellingItems(&sellingItems, queryItemName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error())
	}
	fmt.Println(sellingItems)
	//fmt.Println(sellingItems[0].Name)
	jsonData, err := json.Marshal(sellingItems)
	if err != nil {
		fmt.Println("Error in marshalling")
		return
	}

	// Set Content-Type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to the response
	_, err = w.Write(jsonData)
	fmt.Println(sellingItems)
	if err != nil {
		fmt.Println("Error in writing json to response")
		return
	}
}
