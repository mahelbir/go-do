package system

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

func Database() *sql.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		log.Fatalf("failed to create sqlmock: %s", err)
	}
	return db
}
