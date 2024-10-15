package handler

import (
	"encoding/json"
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
	json.NewDecoder(r.Body).Decode(&user)
	h.userService.AddUser(user)
	w.WriteHeader(http.StatusOK)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if user, err := h.userService.GetUser(name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
