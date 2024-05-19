package services

import (
	"database/sql"
	"errors"
	"go-do/models"
	"time"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	mockUser()
	return &UserService{DB: db}
}

func (r *UserService) GetByID(id int) (*models.User, error) {
	// row := r.DB.QueryRow("SELECT id, full_name, email, password, is_admin, created_at, updated_at FROM users WHERE id = ?", id)

	for _, user := range memUser {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// ============== MOCKING ==============

var memUser []models.User

func mockUser() {
	memUser = append(memUser, models.User{
		ID:        1,
		FullName:  "John Doe",
		Email:     "admin@gmail.com",
		Password:  "$2a$10$RTmgXgIe8LEXQl5Q210KhO1CWLzxz9sAdUPDiDtsvhklbJzr2tq5u",
		IsAdmin:   true,
		CreatedAt: time.Unix(1715483990, 0),
		UpdatedAt: time.Unix(1715483990, 0),
	})
	memUser = append(memUser, models.User{
		ID:        2,
		FullName:  "Jim Doe",
		Email:     "user@gmail.com",
		Password:  "$2a$10$UHbUyJyZ89RcRDJp87zH7eH0pLfe6IxSGkuDO0MApg/ODOxLbTCWy",
		IsAdmin:   false,
		CreatedAt: time.Unix(1715743190, 0),
		UpdatedAt: time.Unix(1715746790, 0),
	})
}
