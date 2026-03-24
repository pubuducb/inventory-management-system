package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const DefaultDBPath = "data/app.db"

func InitDB() (*sql.DB, error) {

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = DefaultDBPath
	}

	db, err := initDB(dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initDB(dbPath string) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbPath)
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(0)

	return db, nil
}

func SetupSchema(db *sql.DB) error {
	if err := createProductsTable(db); err != nil {
		return err
	}
	return nil
}

func createProductsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		archived_at DATETIME DEFAULT NULL
	);`
	_, err := db.Exec(query)
	return err
}
