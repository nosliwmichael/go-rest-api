package repo

import (
	"errors"
	"sync"

	"github.com/nosliwmichael/go-rest-api/internal/model"
)

type UserRepo struct {
	users map[string]*model.User
	mu    sync.Mutex
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[string]*model.User),
	}
}

func (r *UserRepo) AddUser(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.Name] = &user
	return nil
}

func (r *UserRepo) GetUser(name string) (u *model.User, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, exists := r.users[name]
	if !exists {
		err = errors.New("user not found")
	}
	return u, err
}
