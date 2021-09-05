package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pooriaAcademy/event-driven-design-golang/user/internal/core/ports"
	"log"
	"net/http"
)

type HttpHandler struct {
	ports.UserService
}



type UserCreationInput struct {
	UserName string `json:"userName"`
}

func (h HttpHandler) CreateUser(w http.ResponseWriter, req * http.Request){
	input := &UserCreationInput{}
	err := json.NewDecoder(req.Body).Decode(input)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request was sent", http.StatusBadRequest)
		return
	}

	result, err := h.UserService.NewUser(input.UserName)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ToggleFollowInput struct {
	UserId1 string `json:"userId1"`
	UserId2 string `json:"userId2"`
}

func (h HttpHandler) ToggleFollow(w http.ResponseWriter, req * http.Request){
	input := &ToggleFollowInput{}
	err := json.NewDecoder(req.Body).Decode(input)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request was sent", http.StatusBadRequest)
		return
	}

	result, err := h.UserService.ToggleFollowingUser(input.UserId1, input.UserId2)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}






func (h HttpHandler) GetUser(w http.ResponseWriter, req * http.Request){
	vars := mux.Vars(req)

	result, err := h.UserService.GetUser(vars["id"])

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

