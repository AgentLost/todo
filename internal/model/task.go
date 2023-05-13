package model

import "time"

type Task struct {
	Id          int       `json:"id" db:"task_id"`
	UserId      int       `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Due         time.Time `json:"due" db:"due"`
}

type TaskDto struct {
	UserId      int       `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Due         time.Time `json:"due" db:"due"`
}

type TaskRequest struct {
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Due         time.Time `json:"due" db:"due"`
}
