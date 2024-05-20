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

func (s *TodoMessageService) Create(todoMessage models.TodoMessage) (models.TodoMessage, error) {
	// _, err := s.DB.Exec("INSERT INTO todo_message (todo_list_id, content, is_completed) VALUES (?, ?, ?)", todoMessage.TodoListID, todoMessage.Content, todoMessage.IsCompleted)

	todoMessage.ID = memTodoMessage[len(memTodoMessage)-1].ID + 1
	todoMessage.CreatedAt = time.Now()
	todoMessage.UpdatedAt = time.Now()
	todoMessage.DeletedAt = nil
	todoMessage.IsCompleted = false

	memTodoMessage = append(memTodoMessage, todoMessage)
	return todoMessage, nil
}

func (s *TodoMessageService) Update(todoMessage models.TodoMessage) (models.TodoMessage, error) {
	// _, err := s.DB.Exec("UPDATE todo_message SET content = ?, is_completed = ? WHERE id = ?", todoMessage.Content, todoMessage.IsCompleted, todoMessage.ID)

	for i, t := range memTodoMessage {
		if t.ID == todoMessage.ID {
			todoMessage.TodoListID = t.TodoListID
			todoMessage.CreatedAt = t.CreatedAt
			todoMessage.UpdatedAt = time.Now()
			todoMessage.DeletedAt = t.DeletedAt
			memTodoMessage[i] = todoMessage
			return todoMessage, nil
		}
	}
	return models.TodoMessage{}, nil
}

func (s *TodoMessageService) Delete(id int) error {
	// _, err := s.DB.Exec("UPDATE todo_message SET deleted_at = ? WHERE id = ?", time.Now(), id)

	for i, t := range memTodoMessage {
		if t.ID == id {
			now := time.Now()
			t.DeletedAt = &now
			memTodoMessage[i] = t
			return nil
		}
	}
	return nil
}

func (s *TodoMessageService) ListByTodoListID(todoListID int) ([]models.TodoMessage, error) {
	// rows, err := s.DB.Query("SELECT id, todo_list_id, created_at, updated_at, deleted_at, content, is_completed FROM todo_message WHERE todo_list_id = ? AND deleted_at IS NULL", todoListID)

	var todoMessages []models.TodoMessage
	for _, t := range memTodoMessage {
		if t.TodoListID == todoListID && t.DeletedAt == nil {
			todoMessages = append(todoMessages, t)
		}
	}
	return todoMessages, nil
}

func (s *TodoMessageService) GetByID(id int) (models.TodoMessage, error) {
	// row := s.DB.QueryRow("SELECT id, todo_list_id, created_at, updated_at, deleted_at, content, is_completed FROM todo_message WHERE id = ? AND deleted_at IS NULL", id)

	for _, t := range memTodoMessage {
		if t.ID == id && t.DeletedAt == nil {
			return t, nil
		}
	}
	return models.TodoMessage{}, nil
}

func (s *TodoMessageService) SetCompleted(id int, isCompleted bool) error {
	// _, err := s.DB.Exec("UPDATE todo_message SET is_completed = ? WHERE id = ?", isCompleted, id)

	for i, t := range memTodoMessage {
		if t.ID == id {
			t.IsCompleted = isCompleted
			memTodoMessage[i] = t
			return nil
		}

	}

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
