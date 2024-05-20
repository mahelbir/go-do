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

func (s *TodoListService) Create(todoList models.TodoList) (models.TodoList, error) {
	// _, err := s.DB.Exec("INSERT INTO todo_list (user_id, title) VALUES (?, ?)", todoList.UserID, todoList.Title)

	todoList.ID = memTodoList[len(memTodoList)-1].ID + 1
	todoList.CreatedAt = time.Now()
	todoList.UpdatedAt = time.Now()

	memTodoList = append(memTodoList, todoList)
	return todoList, nil
}

func (s *TodoListService) Update(todoList models.TodoList) (models.TodoList, error) {
	// _, err := s.DB.Exec("UPDATE todo_list SET title = ? WHERE id = ?", todoList.Title, todoList.ID)

	for i, t := range memTodoList {
		if t.ID == todoList.ID {
			todoList.UpdatedAt = time.Now()
			memTodoList[i] = todoList
			return todoList, nil
		}
	}
	return models.TodoList{}, nil
}

func (s *TodoListService) Delete(id int) error {
	// _, err := s.DB.Exec("DELETE FROM todo_list WHERE id = ?", id)

	for i, t := range memTodoList {
		if t.ID == id {
			memTodoList = append(memTodoList[:i], memTodoList[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *TodoListService) List() ([]models.TodoList, error) {
	// rows, err := s.DB.Query("SELECT * FROM todo_list")

	return memTodoList, nil
}

func (s *TodoListService) ListByUserID(userID int) ([]models.TodoList, error) {
	// rows, err := s.DB.Query("SELECT * FROM todo_list WHERE user_id = ?", userID)

	todoLists := make([]models.TodoList, 0)
	for _, todoList := range memTodoList {
		if todoList.UserID == userID {
			todoLists = append(todoLists, todoList)
		}
	}
	return todoLists, nil
}

func (s *TodoListService) GetByID(id int) (models.TodoList, error) {
	// row := s.DB.QueryRow("SELECT * FROM todo_list WHERE id = ?", id)

	for _, todoList := range memTodoList {
		if todoList.ID == id {
			return todoList, nil
		}
	}
	return models.TodoList{}, nil
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
