package main

import (
	"github.com/gorilla/mux"
	"github.com/pooriaAcademy/event-driven-design-golang/user"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/user/follow", user.HttpHandler.ToggleFollow).Methods("POST")
	r.HandleFunc("/user", user.HttpHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", user.HttpHandler.GetUser).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	log.Fatalln(err)
}
