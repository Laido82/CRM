package controllers

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() (*sql.DB, error) {

	// Get absolute path to the database directory
	filepathAbs, err := GetAbsPath("../internal/database/")
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path: %w", err)
	}

	databaseConfig, err := GetDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading database config: %w", err)
	}

	driver := databaseConfig.DBDriver
	dbName := databaseConfig.DBName
	db, err := sql.Open(driver, filepath.Join(filepathAbs, dbName))

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}

func CloseDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}
	return nil
}
