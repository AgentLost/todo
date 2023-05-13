package service_impl

import (
	"log"
	"todo-app/internal/helpers"
	"todo-app/internal/model"
	"todo-app/internal/repo"
)

type User struct {
	repo repo.UserRepository
}

func (u *User) Register(user model.UserRequest) error {
	log.Printf("registred user %s", user)

	newPassword, err := helpers.Encode(user.Password)
	user.Password = newPassword

	if err != nil {
		return err
	}

	return u.repo.Save(user)
}

func (u *User) Update(userId int, user model.UserRequest) error {
	log.Printf("update user %s", user)

	newPassword, err := helpers.Encode(user.Password)
	user.Password = newPassword

	if err != nil {
		return err
	}

	return u.repo.Update(userId, user)
}

func (u *User) Delete(userId int) error {
	log.Printf("delete user with id = %d", userId)
	err := u.repo.Delete(userId)

	return err
}

func (u *User) GetByEmail(email string) (user model.User, err error) {
	log.Printf("get user with email = %s", email)

	user, err = u.repo.GetByEmail(email)

	log.Printf("get user %v", user)
	return
}

func (u *User) Get(userId int) (user model.User, err error) {
	log.Printf("get user with id = %d", userId)

	user, err = u.repo.GetById(userId)

	log.Printf("get user %v", user)
	return
}

func NewUserService(repo repo.UserRepository) *User {
	return &User{repo: repo}
}
