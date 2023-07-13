package database

import (
	_ "database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var database *sqlx.DB

func GetDB() *sqlx.DB {
	if database == nil {
		LoadDatabase()
	}
	return database
}

func LoadDatabase() {
	// Connection string Tt8UZj-sXj6G6cLkKUaBQA
	connString := "postgres://postgres:admin@localhost:5432/web_coffee"

	// Create a new DB object
	db, err := sqlx.Open("pgx", connString)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	database = db
}
