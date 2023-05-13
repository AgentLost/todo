package repo

import (
	"todo-app/internal/infrastructure/database"
	"todo-app/internal/model"
	"todo-app/internal/repo/postgres"
)

type Repository struct {
	UserRepository
	TaskRepository
}

type UserRepository interface {
	Save(user model.UserRequest) error
	Update(userId int, user model.UserRequest) error
	Delete(userId int) error
	GetByEmail(email string) (model.User, error)
	GetById(userId int) (model.User, error)
}

type TaskRepository interface {
	Save(task model.TaskDto) error
	Update(taskId int, task model.TaskDto) error
	Delete(taskId int) error
	GetById(taskId int) (model.Task, error)
	GetAllByUserId(userId int) ([]model.Task, error)
}

func New(store *database.Store) *Repository {
	return &Repository{postgres.NewUserRepository(store), postgres.NewTaskRepository(store)}
}
