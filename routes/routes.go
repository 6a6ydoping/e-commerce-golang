package routes

import (
	"e-commerce-app/controllers"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func RegisterRoutes() {
	Router = mux.NewRouter()
	Router.HandleFunc("/", controllers.HandleRegistration)
	Router.HandleFunc("/register", controllers.HandleRegistration)
	Router.HandleFunc("/login", controllers.HandleLogin)
	Router.HandleFunc("/auth", controllers.HandleLogin)
	Router.HandleFunc("/home", controllers.Home)
	Router.HandleFunc("/createItem", controllers.CreateItem)
	Router.HandleFunc("/menu", controllers.GetAllSellingItems)
}
