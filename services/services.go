package services

import (
	"database/sql"
)

var (
	UserSvc        *UserService
	TodoListSvc    *TodoListService
	TodoMessageSvc *TodoMessageService
)

func InitServices(db *sql.DB) {
	UserSvc = NewUserService(db)
	TodoListSvc = NewTodoListService(db)
	TodoMessageSvc = NewTodoMessageService(db)
}
