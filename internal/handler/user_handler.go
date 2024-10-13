package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nosliwmichael/go-rest-api/internal/model"
	"github.com/nosliwmichael/go-rest-api/internal/service"
)

type (
	UserHandler interface {
		AddUser(w http.ResponseWriter, r *http.Request)
		GetUser(w http.ResponseWriter, r *http.Request)
	}
	userHandler struct {
		userService service.UserService
	}
)

var _ UserHandler = userHandler{}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (h userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	h.userService.AddUser(user)
	w.WriteHeader(http.StatusOK)
}

func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if user, err := h.userService.GetUser(name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
