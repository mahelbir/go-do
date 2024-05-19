package services

import (
	"database/sql"
	"go-do/models"
	"time"
)

type TodoListService struct {
	DB *sql.DB
}

func NewTodoListService(db *sql.DB) *TodoListService {
	mockTodoList()
	return &TodoListService{DB: db}
}

// ============== MOCKING ==============

var memTodoList []models.TodoList

func mockTodoList() {
	memTodoList = append(memTodoList, models.TodoList{
		ID:        1,
		UserID:    2,
		CreatedAt: time.Unix(1716089965, 0),
		UpdatedAt: time.Unix(1716089965, 0),
		DeletedAt: nil,
		Title:     "Title 1",
	})
	memTodoList = append(memTodoList, models.TodoList{
		ID:        2,
		UserID:    2,
		CreatedAt: time.Unix(1716090019, 0),
		UpdatedAt: time.Unix(1716090019, 0),
		DeletedAt: nil,
		Title:     "Title 2",
	})
	memTodoList = append(memTodoList, models.TodoList{
		ID:        3,
		UserID:    1,
		CreatedAt: time.Unix(1716090370, 0),
		UpdatedAt: time.Unix(1716090370, 0),
		DeletedAt: nil,
		Title:     "Title 3",
	})
}
