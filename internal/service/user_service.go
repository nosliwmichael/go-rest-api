package service

import (
	"github.com/nosliwmichael/go-rest-api/internal/model"
	"github.com/nosliwmichael/go-rest-api/internal/repo"
)

type (
	UserService interface {
		AddUser(model.User) error
		GetUser(string) (*model.User, error)
	}
	userService struct {
		userRepo repo.UserRepo
	}
)

var _ UserService = userService{}

func NewUserService(userRepo repo.UserRepo) userService {
	return userService{
		userRepo: userRepo,
	}
}

func (r userService) AddUser(user model.User) error {
	return r.userRepo.AddUser(user)
}

func (r userService) GetUser(name string) (*model.User, error) {
	return r.userRepo.GetUser(name)
}
