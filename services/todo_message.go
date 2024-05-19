package services

import (
	"database/sql"
	"go-do/models"
	"time"
)

type TodoMessageService struct {
	DB *sql.DB
}

func NewTodoMessageService(db *sql.DB) *TodoMessageService {
	mockTodoMessage()
	return &TodoMessageService{DB: db}
}

func (r *TodoMessageService) GetById(id int) (*models.TodoMessage, error) {
	// row := r.DB.QueryRow("SELECT id, todo_list_id, created_at, updated_at, deleted_at, content, is_completed FROM todo_messages WHERE id = ?", id)

	todoMessage := &models.TodoMessage{}
	return todoMessage, nil
}

func (r *TodoMessageService) Create(todoMessage *models.TodoMessage) error {
	// _, err := r.DB.Exec("INSERT INTO todo_messages (todo_list_id, created_at, updated_at, content, is_completed) VALUES (?, ?, ?, ?, ?)", todoMessage.TodoListID, todoMessage.CreatedAt, todoMessage.UpdatedAt, todoMessage.Content, todoMessage.IsCompleted)

	return nil
}

func (r *TodoMessageService) Update(todoMessage *models.TodoMessage) error {
	// _, err := r.DB.Exec("UPDATE todo_messages SET updated_at = ?, deleted_at = ?, content = ?, is_completed = ? WHERE id = ?", todoMessage.UpdatedAt, todoMessage.DeletedAt, todoMessage.Content, todoMessage.IsCompleted, todoMessage.ID)

	return nil
}

func (r *TodoMessageService) Delete(id int) error {
	// _, err := r.DB.Exec("DELETE FROM todo_messages WHERE id = ?", id)

	return nil
}

// ============== MOCKING ==============

var memTodoMessage []models.TodoMessage

func mockTodoMessage() {
	memTodoMessage = append(memTodoMessage, models.TodoMessage{
		ID:          1,
		TodoListID:  1,
		CreatedAt:   time.Unix(1716089966, 0),
		UpdatedAt:   time.Unix(1716089966, 0),
		DeletedAt:   nil,
		Content:     "Item 1",
		IsCompleted: false,
	})
	memTodoMessage = append(memTodoMessage, models.TodoMessage{
		ID:          2,
		TodoListID:  1,
		CreatedAt:   time.Unix(1716090020, 0),
		UpdatedAt:   time.Unix(1716090020, 0),
		DeletedAt:   nil,
		Content:     "Item 2",
		IsCompleted: true,
	})
	memTodoMessage = append(memTodoMessage, models.TodoMessage{
		ID:          3,
		TodoListID:  2,
		CreatedAt:   time.Unix(1716090032, 0),
		UpdatedAt:   time.Unix(1716090032, 0),
		DeletedAt:   nil,
		Content:     "Item 1",
		IsCompleted: true,
	})
	memTodoMessage = append(memTodoMessage, models.TodoMessage{
		ID:          4,
		TodoListID:  3,
		CreatedAt:   time.Unix(1716090371, 0),
		UpdatedAt:   time.Unix(1716090371, 0),
		DeletedAt:   nil,
		Content:     "Item 1",
		IsCompleted: false,
	})
}
