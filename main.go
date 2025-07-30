package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"restapitry/config"
	_ "restapitry/docs"
	"restapitry/routers"
)

// @title Subscriptions API
// @version 1.0
// @description REST-сервис для учёта онлайн-подписок
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectDB()
	routes := routers.Routes()
	routes.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", routes))

}
