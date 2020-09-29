package app

import (
	"github.com/nickdipreta/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp(){
	//initialize the handlers that we have "/users" controller is user controller, function used is GetUser
	http.HandleFunc("/users", controllers.GetUser)
	// launch application with ListenAndServe
	if err := http.ListenAndServe(":8080", nil); err!=nil {
		panic(err)
	}
}