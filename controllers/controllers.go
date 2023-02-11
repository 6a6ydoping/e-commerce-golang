package controllers

import (
	"e-commerce-app/initializers"
	"e-commerce-app/middlewares"
	"e-commerce-app/models"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
		fmt.Println(email, firstName, userType, reflect.TypeOf(password))

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

func CreateItem(w http.ResponseWriter, r *http.Request) {
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

func GetAllSellingItems(w http.ResponseWriter, r *http.Request) {
	var sellingItems []models.Item
	if err := initializers.DB.Select("id, name, price, quantity, rating").Find(&sellingItems).Error; err != nil {
		fmt.Println("Failed to fetch all selling items")
		http.Error(w, "Failed to fetch all selling items", http.StatusInternalServerError)
	}
	jsonData, err := json.Marshal(sellingItems)
	if err != nil {
		// Handle error
		return
	}

	// Set Content-Type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		// Handle error
		return
	}
}
