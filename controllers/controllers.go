package controllers

import (
	"e-commerce-app/initializers"
	"e-commerce-app/middlewares"
	"e-commerce-app/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Берем токен
		token, err := middlewares.GetTokenValueFromCookie(r)
		if err != nil {
			http.Error(w, "You have to login!", http.StatusUnauthorized)
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
	err := initializers.DB.Find(&sellingItems).Error
	if err != nil {
		fmt.Println("Failed to fetch all selling items")
		http.Error(w, "Failed to fetch all selling items", http.StatusInternalServerError)
	}
	fmt.Println(sellingItems)
}
