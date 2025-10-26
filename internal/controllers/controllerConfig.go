package controllers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	DBDriver string
	Username string
	Password string
	DBName   string
	Host     string
	Port     string
}

func loadEnv() {
	// Get absolute path of the .env file and load it
	absPath, err := GetAbsPath("../internal/env/")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	err = godotenv.Load(filepath.Join(absPath, ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetServerConfig() (*ServerConfig, error) {
	loadEnv()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return &ServerConfig{
		Host: host,
		Port: port,
	}, nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	loadEnv()

	driver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	return &DatabaseConfig{
		DBDriver: driver,
		Username: username,
		Password: password,
		DBName:   dbName,
		Host:     host,
		Port:     port,
	}, nil
}
