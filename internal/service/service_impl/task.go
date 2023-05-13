package service_impl

import (
	"errors"
	"log"
	"todo-app/internal/model"
	"todo-app/internal/repo"
)

type Task struct {
	repo repo.TaskRepository
}

func (t *Task) Save(task model.TaskDto) error {
	log.Printf("save task %v", task)

	return t.repo.Save(task)
}

func (t *Task) Update(userId int, taskId int, task model.TaskDto) error {
	log.Printf("update task with id = %d task: %v", taskId, task)

	before, err := t.repo.GetById(taskId)

	if err != nil {
		return err
	}

	if before.UserId != userId || userId != task.UserId {
		return errors.New("ids not matches")
	}

	return t.repo.Update(taskId, task)
}

func (t *Task) Delete(userId int, taskId int) error {
	log.Printf("delete task with taskId = %d", taskId)

	task, err := t.repo.GetById(taskId)

	if err != nil {
		return err
	}

	if task.UserId != userId {
		return errors.New("ids not matches")
	}

	return t.repo.Delete(taskId)
}

func (t *Task) Get(userId int, taskId int) (model.Task, error) {
	log.Printf("get task by id = %d", taskId)

	task, err := t.repo.GetById(taskId)

	if err != nil {
		return model.Task{}, err
	}

	if task.UserId != userId {
		return model.Task{}, errors.New("ids not matches")
	}

	return task, err
}

func (t *Task) GetAll(userId int) ([]model.Task, error) {
	log.Printf("get tasks by user id = %d", userId)
	tasks, err := t.repo.GetAllByUserId(userId)

	return tasks, err
}

func NewTaskService(repo repo.TaskRepository) *Task {
	return &Task{repo: repo}
}
