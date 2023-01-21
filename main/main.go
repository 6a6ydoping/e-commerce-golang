package main

import (
	"e-commerce-app/initializers"
	"e-commerce-app/routes"
	"fmt"
	"net/http"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDataBase()
	initializers.SyncDB()
	routes.RegisterRoutes()
}

func main() {
	fmt.Println("Successfully running on :8000 port!!")
	http.ListenAndServe("localhost:8000", routes.Router)
}
