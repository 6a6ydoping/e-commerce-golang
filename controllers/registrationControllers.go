package controllers

import (
	"e-commerce-app/helpers"
	"e-commerce-app/middlewares"
	"e-commerce-app/models"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w)
	// перешли по url/register -> грузим вьюшку
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/register.html")
		if err != nil {
			log.Fatal("Error parsing file...")
		}
		t.Execute(w, nil)

	} else if r.Method == "POST" { //отправили запрос на регистрацию
		var requestBody map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Body == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		r.ParseForm()

		//парсинг формы
		firstName := fmt.Sprintf("%v", requestBody["firstName"])
		lastName := fmt.Sprintf("%v", requestBody["lastName"])
		email := fmt.Sprintf("%v", requestBody["email"])
		password := fmt.Sprintf("%v", requestBody["password"])
		userType := fmt.Sprintf("%v", requestBody["userType"])

		//хешируем пароль
		passwordHash, err := middlewares.HashPassword(password)
		if err != nil {
			log.Fatal("error in hashing pass")
		}

		//проверка есть ли в бд пользователь с таким имейлом
		if err := middlewares.CheckEmailInDB(email, userType); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//загружаем в бд
		switch userType {
		case "Seller":
			err = middlewares.InsertSellerIntoDataBase(models.CreateSeller(firstName, lastName, email, passwordHash))
		case "Client":
			err = middlewares.InsertClientIntoDataBase(models.CreateClient(firstName, lastName, email, passwordHash))
		default:
			log.Fatal("Unknown user type")
		}
		if err != nil {
			log.Fatal("Error saving user to db")
		}

		//Редиректим на url/login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println("User successfully added to db")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		helpers.EnableCors(&w)

		var requestBody map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(requestBody) == 0 {
			http.Error(w, "empty request body", http.StatusBadRequest)
			return
		}
		fmt.Println(requestBody)
		//парсим форму логина
		email := fmt.Sprintf("%v", requestBody["email"])
		password := fmt.Sprintf("%v", requestBody["password"])
		userType := fmt.Sprintf("%v", requestBody["userType"])

		//ищем юзера по почте
		err = middlewares.CheckEmailAndPasswordInDB(email, password, userType)
		if err != nil {
			http.Error(w, "No such user", http.StatusBadRequest)
		}
		//Создать токен и создать реквест с токеном
		switch userType {
		case "Seller":
			var seller, err = middlewares.GetSellerByEmail(email)
			if err != nil {
				http.Error(w, "No such user", http.StatusBadRequest)
			}
			token, err := middlewares.CreateSellerToken(seller)
			if err != nil {
				http.Error(w, "Cant create jwt token", http.StatusBadRequest)
			}
			middlewares.DeleteTokenFromCookie(w, r)
			middlewares.SetTokenToCookie(w, token)
		case "Client":
			var client, err = middlewares.GetClientByEmail(email)
			if err != nil {
				http.Error(w, "No such user", http.StatusBadRequest)
			}
			token, err := middlewares.CreateClientToken(client)
			if err != nil {
				http.Error(w, "Cant create jwt token", http.StatusBadRequest)
			}
			middlewares.DeleteTokenFromCookie(w, r)
			middlewares.SetTokenToCookie(w, token)
		}

		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
