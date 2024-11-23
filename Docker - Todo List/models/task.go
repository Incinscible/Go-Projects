package models

import "time"

type Task struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Completed   bool      `db:"completed" json:"completed"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
