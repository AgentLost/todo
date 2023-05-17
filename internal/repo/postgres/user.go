package postgres

import (
	"todo-app/internal/infrastructure/database"
	"todo-app/internal/model"
)

type User struct {
	store *database.Postgres
}

func (u *User) Save(user model.UserRequest) error {
	_, err := u.store.DB.Exec(
		"insert into users (username, email, password) values ($1, $2, $3)",
		user.Username, user.Email, user.Password)
	return err
}

func (u *User) Update(userId int, user model.UserRequest) error {
	_, err := u.store.DB.Exec(
		"update users set username=$1, email=$2, password=$3 where user_id=$4",
		user.Username, user.Email, user.Password, userId)

	return err
}

func (u *User) Delete(userId int) error {
	_, err := u.store.DB.Exec("delete from users where user_id=$1", userId)

	return err
}

func (u *User) GetByEmail(email string) (model.User, error) {
	user := model.User{}

	err := u.store.DB.Get(&user, "select * from users where email=$1", email)

	return user, err
}

func (u *User) GetById(userId int) (model.User, error) {
	user := model.User{}

	err := u.store.DB.Get(&user, "select * from users where user_id=$1", userId)

	return user, err
}

func NewUserRepository(store *database.Postgres) *User {
	return &User{store: store}
}
