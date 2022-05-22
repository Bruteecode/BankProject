package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	RegisterBankingDetailsRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
func RegisterBankingDetailsRoutes(router *mux.Router) {
	router.HandleFunc("/api/Create", controllers.Create).Methods("POST")
	router.HandleFunc("/api/Add", controllers.Add).Methods("PUT")
	router.HandleFunc("/api/Deduct", controllers.Deduct).Methods("PUT")
	router.HandleFunc("/api/Close", controllers.Close).Methods("DELETE")
	router.HandleFunc("/api/GetInformationBy", controllers.GetBalance).Methods("GET")
}
