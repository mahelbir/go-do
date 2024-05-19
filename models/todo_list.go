package models

import "time"

type TodoList struct {
	ID             int        `json:"id"`
	UserID         int        `json:"user_id"`
	Title          string     `json:"title"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	CompletionRate int        `json:"completion_rate"`
}
