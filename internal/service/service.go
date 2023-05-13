package service

import (
	"todo-app/internal/helpers"
	"todo-app/internal/model"
	"todo-app/internal/repo"
	"todo-app/internal/service/service_impl"
)

type UserService interface {
	Register(user model.UserRequest) error
	Update(userId int, user model.UserRequest) error
	Delete(userId int) error
	GetByEmail(email string) (model.User, error)
	Get(userId int) (model.User, error)
}

type TaskService interface {
	Save(task model.TaskDto) error
	Update(userId int, taskId int, task model.TaskDto) error
	Delete(userId int, taskId int) error
	Get(userId int, taskId int) (model.Task, error)
	GetAll(userId int) ([]model.Task, error)
}

type AuthService interface {
	Validate(token string) (bool, error)
	GetHeader() string
	Generate(userId int) (string, error)
	GetUserId(token string) (int, error)
}

type Service struct {
	UserService
	TaskService
	AuthService
}

func New(r *repo.Repository, p *helpers.TokenProvider) *Service {
	return &Service{
		service_impl.NewUserService(r.UserRepository),
		service_impl.NewTaskService(r.TaskRepository),
		service_impl.NewAuthService(p),
	}
}
