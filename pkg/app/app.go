package app

import (
	"fmt"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/domain/repo"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/service"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"

)

func Start() {
	router := mux.NewRouter()
	dbClient := GetMongoClient()

	dbName := "userdb"
	timeout := time.Second * 5

	//User Wiring
	userRepo := repo.NewUserRepo(dbClient, dbName, timeout)
	userService := service.NewUserService(userRepo)
	userHandler := UserHandler{service: userService}

	//Phone Wiring
	phoneService := service.NewPhoneService()
	phoneHandler := PhoneHandler{service: phoneService}

	//User handler mapping
	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/available", userHandler.CheckAvailable).Queries("username", "{username}").Methods(http.MethodGet)
	router.HandleFunc("/users/available", userHandler.CheckAvailable).Queries("email", "{email}").Methods(http.MethodGet)
	router.HandleFunc("/users/available", userHandler.CheckAvailable).Queries("phone", "{phone}").Methods(http.MethodGet)

	//Phone handler mapping
	router.HandleFunc("/phone/otp", phoneHandler.PhoneOtp).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s", "localhost:8080"), router))

}



