package repo

import (
	"database/sql"
	"github.com/nosliwmichael/go-rest-api/internal/model"
	"github.com/nosliwmichael/go-rest-api/internal/repo/queries"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{db: db}
}

func (r UserRepo) AddUser(user model.User) error {
	_, err := r.db.Exec(queries.InsertUser, user.Username, user.FirstName, user.LastName, user.Age)
	return err
}

func (r UserRepo) GetUser(id string) (u *model.User, err error) {
	row := r.db.QueryRow(queries.SelectUserById, id)
	var user model.User
	err = row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Age)
	u = &user
	return
}
