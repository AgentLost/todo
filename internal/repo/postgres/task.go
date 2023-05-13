package postgres

import (
	"todo-app/internal/infrastructure/database"
	"todo-app/internal/model"
)

type Task struct {
	store *database.Store
}

func (t *Task) GetAllByUserId(userId int) ([]model.Task, error) {
	var tasks []model.Task

	err := t.store.DB.Select(&tasks, "select * from task where user_id=$1", userId)

	return tasks, err
}

func (t *Task) Save(task model.TaskDto) error {
	_, err := t.store.DB.Exec(
		"insert into task (user_id, title, description, due) values ($1,$2,$3,$4)",
		task.UserId, task.Title, task.Description, task.Due)

	return err
}

func (t *Task) Update(taskId int, task model.TaskDto) error {
	_, err := t.store.DB.Exec(
		"update task set user_id=$1, title=$2, description=$3, due=$4 where task_id=$5",
		task.UserId, task.Title, task.Description, task.Due, taskId)

	return err
}

func (t *Task) Delete(taskId int) error {
	_, err := t.store.DB.Exec("delete from task where task_id=$1", taskId)
	return err
}

func (t *Task) GetById(taskId int) (model.Task, error) {
	task := model.Task{}

	err := t.store.DB.Get(&task, "select * from task where task_id=$1", taskId)

	return task, err
}

func NewTaskRepository(store *database.Store) *Task {
	return &Task{store: store}
}
