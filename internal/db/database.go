package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	runMigrations()
}

func runMigrations() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		done BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed_at DATETIME
	);`

	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Erro ao criar tabela de tarefas: %v", err)
	}
}
