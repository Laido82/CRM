package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func MakeCopyDB() error {

	// Get absolute path to the database directory

	filepathAbs, err := GetAbsPath("../internal/")
	if err != nil {
		return fmt.Errorf("error getting absolute path: %w", err)
	}
	// Define the copy database name
	copyDBName := time.Now().Format("20060102") + "_contacts_copy.db"

	//Save the copy in different directory
	copyDBPath := filepath.Join(filepathAbs, "backups", copyDBName)

	// Original database path
	originalDBPath := filepath.Join(filepathAbs, "database", "contacts.db")
	//Do the copy
	originalDB, err := os.ReadFile(originalDBPath)
	if err != nil {
		return fmt.Errorf("error reading original database: %w", err)
	}

	// Compress and write the database copy
	err = os.WriteFile(copyDBPath, originalDB, 0644)
	if err != nil {
		return fmt.Errorf("error writing copy database: %w", err)
	}
	return nil
}
