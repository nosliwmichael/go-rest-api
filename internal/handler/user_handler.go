package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/nosliwmichael/go-rest-api/internal/model"
)

type (
	UserService interface {
		AddUser(model.User) error
		GetUser(string) (*model.User, error)
	}
	UserHandler struct {
		userService UserService
	}
)

func NewUserHandler(userService UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
	}
	err = h.userService.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if user, err := h.userService.GetUser(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
	} else {
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to write body: %+v", err)
		}
	}
}
