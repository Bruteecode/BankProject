package controllers

import (
	"encoding/json"
	"fmt"
	"golang-crud-rest-api/entities"

	// "golang-crud-rest-api/database"
	// "golang-crud-rest-api/entities"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// "github.com/example/simple-REST/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Connector *gorm.DB

// func Connect(connectionString string) error {
// 	var err error
// 	Connector, err = gorm.Open("mysql", ConnectionString)
// 	if err != nil {
// 		return err

// 	}
// 	log.Println("Connection was successful!!")
// 	return nil
// }

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var Information entities.Information
	json.NewDecoder(r.Body).Decode(&Information)
	Instance.Create(&Information)
	json.NewEncoder(w).Encode(Information)
	fmt.Println(Information)
}
func Add(w http.ResponseWriter, r *http.Request) {
	InformationId := mux.Vars(r)["Id"]

	// if CheckIfInformationExists(InformationId) == false {
	// 	json.NewEncoder(w).Encode("Account not found!")
	// 	return
	// }
	var Information entities.Information
	Instance.First(&Information, InformationId)
	json.NewDecoder(r.Body).Decode(&Information)
	Instance.Model(&entities.Information{}).Where("AccountNumber=?", Information.AccountNumber).Updates(entities.Information{AccountBalance: Information.AccountBalance})

	fmt.Fprintf(w, "AccountBalance:%+v", Information.AccountBalance)

	// database.Instance.Update(&Information)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Information)

}
func Deduct(w http.ResponseWriter, r *http.Request) {
	InformationId := mux.Vars(r)["Id"]
	// if checkIfInformationExists(InformationId) == false {
	// if checkIfInformationExists(InformationId) == false {

	// 	json.NewEncoder(w).Encode("Account not found")
	// 	return
	// }
	var Information entities.Information
	Instance.First(&Information, InformationId)
	json.NewDecoder(r.Body).Decode(&Information)
	Instance.Save(&Information)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Information)
}
func Close(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	InformationId := mux.Vars(r)["id"]
	// if checkIfInformationExists(InformationId) == false {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode("Account not found")
	// 	return
	// }
	var Information entities.Information
	Instance.Delete(&Information, InformationId)
	json.NewEncoder(w).Encode("Account Closed Succesfully")
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	// InformationId := mux.Vars(r)["Id"]
	var Information entities.Information
	// database.Instance.First(&Information, InformationId)
	json.NewDecoder(r.Body).Decode(&Information)
	Instance.Order("Transaction _time desc").Limit(1).Find(&Information)
	// database.Limit(1).Find(&Information)
	// database.Instance.Model(&entities.Information{}).Where("AccountNumber=?", Information.AccountNumber).Updates(entities.Information{TransactionTime: Information.TransactionTime})
	// fmt.Fprintf(w, "TransactionTime:%+v", Information.TransactionTime)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Information)

}

// func getbalance
// select * from Information
// where accountnumber=
// Order by transaction time descending limit 1
// db.Order("age desc, name").Find(&users)
// // SELECT * FROM users ORDER BY age desc, name;

// // Multiple orders
// db.Order("age desc").Order("name").Find(&users)
// // SELECT * FROM users ORDER BY age desc, name;
// db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// db.order().limit().Find()
