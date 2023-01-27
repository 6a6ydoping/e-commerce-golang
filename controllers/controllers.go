package controllers

import (
	"e-commerce-app/middlewares"
	"e-commerce-app/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	// перешли по url/register -> грузим вьюшку
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/register.html")
		if err != nil {
			log.Fatal("Error parsing file...")
		}
		t.Execute(w, nil)

	} else if r.Method == "POST" { //отправили запрос на регистрацию
		r.ParseForm()

		//парсинг формы
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		password := r.FormValue("password")
		userType := r.FormValue("userType")

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

		//JWT???
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println("User successfully added to db")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Fatal("Error parsing file...")
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		//парсим форму логина
		email, password, userType := r.FormValue("email"), r.FormValue("password"), r.FormValue("userType")

		//ищем юзера по почте
		err := middlewares.CheckEmailAndPasswordInDB(email, password, userType)
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
			middlewares.SetTokenToCookie(w, token)
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("views/home.html")
		if err != nil {
			log.Fatal("Error parsing file...")
		}
		t.Execute(w, nil)
	}
}

//func CheckCookie(w http.ResponseWriter, r *http.Request) {
//	token, err := r.Cookie("token")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(token)
//}
