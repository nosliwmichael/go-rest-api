package service

import (
	"github.com/nosliwmichael/go-rest-api/internal/model"
)

type (
	UserRepo interface {
		AddUser(model.User) error
		GetUser(string) (*model.User, error)
	}
	UserService struct {
		userRepo UserRepo
	}
)

func NewUserService(userRepo UserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (r UserService) AddUser(user model.User) error {
	return r.userRepo.AddUser(user)
}

func (r UserService) GetUser(name string) (*model.User, error) {
	return r.userRepo.GetUser(name)
}
