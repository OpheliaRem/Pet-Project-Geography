package main

import (
	"backPet0/db"
	"backPet0/logger"
	"backPet0/logic/country_controller"
	"backPet0/logic/town_contoller"
	"log"
	"net/http"
	"os"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/v1/save-country":      country_controller.Save,
	"/v1/get-country-by-id": country_controller.GetById,
	"/v1/get-all-countries": country_controller.GetAll,
	"/v1/remove-country":    country_controller.Remove,
	"/v1/update-country":    country_controller.Update,
	"/v1/save-town":         town_contoller.Save,
	"/v1/get-all-towns":     town_contoller.GetAll,
	"/v1/get-town-by-id":    town_contoller.GetById,
}

func main() {
	logger.AddFileStream(os.Stdout)

	db.OpenConnection(db.Config{
		Username: "postgres",
		Password: "postgres",
		DbName:   "geography",
	})
	defer db.CloseConnection()

	for endpoint, handler := range handlers {
		http.HandleFunc(endpoint, handler)
	}

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
