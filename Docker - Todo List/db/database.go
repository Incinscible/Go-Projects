package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	if DB != nil {
		log.Println("Database already initialized")
		return
	}
	host := os.Getenv("DB_HOST")
	port := 5432
	user := "user"
	password := os.Getenv("DB_PASSWORD")
	dbname := "todolist"
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to the Database!")
	}

	createTableIfNotExists(DB)
}

func createTableIfNotExists(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		description TEXT,
		completed BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT NOW()
	)`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}
