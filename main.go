package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/henryhiraki22/lifts/controller"
	"github.com/henryhiraki22/lifts/database"
	"github.com/henryhiraki22/lifts/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/api", controller.HealthCheck).Methods("GET")
	router.HandleFunc("/api/createlifta", controller.CreateLiftA).Methods("POST")
	router.HandleFunc("/api/getlifta", controller.ListTrainingA).Methods("GET")
	router.HandleFunc("/api/createliftb", controller.CreateLiftB).Methods("POST")
	router.HandleFunc("/api/getliftb", controller.ListTrainingB).Methods("GET")
}

func initDB() {
	config :=
		database.Config{
			Host:         os.Getenv("MYSQL_HOST"),
			UserName:     os.Getenv("MYSQL_USER"),
			Password:     os.Getenv("MYSQL_PASSWORD"),
			DatabaseName: os.Getenv("MYSQL_DATABASE"),
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.LiftA{}, &entity.LiftB{})
}
