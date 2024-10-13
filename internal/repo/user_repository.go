package repo

import (
	"errors"
	"sync"

	"github.com/nosliwmichael/go-rest-api/internal/model"
)

type (
	UserRepo interface {
		AddUser(model.User) error
		GetUser(string) (*model.User, error)
	}
	userRepo struct {
		users map[string]*model.User
		mu    sync.Mutex
	}
)

var _ UserRepo = (*userRepo)(nil)

func NewUserRepo() *userRepo {
	return &userRepo{
		users: make(map[string]*model.User),
	}
}

func (r *userRepo) AddUser(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.Name] = &user
	return nil
}

func (r *userRepo) GetUser(name string) (u *model.User, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, exists := r.users[name]
	if !exists {
		err = errors.New("user not found")
	}
	return u, err
}
