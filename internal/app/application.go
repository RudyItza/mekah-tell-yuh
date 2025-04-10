package app

import (
	"database/sql"
)

type Application struct {
	DB *sql.DB
}

func NewApplication(db *sql.DB) *Application {
	return &Application{DB: db}
}
