package main

import (
	"log"
	"main/internal/controllers"
	"main/internal/handlers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect to the database
	db, err := controllers.ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	// Ensure the database connection is closed when main exits
	defer func() {
		if err := controllers.CloseDB(db); err != nil {
			log.Println("Error closing database:", err)
		}
	}()
	// Create contacts table if it doesn't exist
	err = controllers.CreateContactsTable(db)
	if err != nil {
		log.Fatal("Error creating contacts table:", err)
	}

	// Load environment variables from .env file
	serverConfig, err := controllers.GetServerConfig()
	if err != nil {
		log.Fatal("Error loading server config:", err)
	}

	host := serverConfig.Host
	port := serverConfig.Port

	if host == "" || port == "" {
		log.Fatal("HOST or PORT environment variables are not set")
	}
	// Parse all templates
	templates := handlers.ParseAllTemplates()

	// Initialize Fiber app
	app := fiber.New()

	// Serve static file
	filepathAbs, err := controllers.GetAbsPath("../static")
	if err != nil {
		log.Fatal("Error getting absolute path for static files:", err)
	}
	app.Static("/static", filepathAbs)

	// Define route handlers
	handlers.Menu(app, templates)
	handlers.AddContact(app, templates, db)
	handlers.UpdateContact(app, templates, db)
	handlers.DeleteContact(app, templates, db)
	handlers.HandlerServices(app, templates)
	handlers.About(app, templates)
	handlers.GetAllContacts(app, templates, db)
	handlers.ListSearchedContacts(app, templates, db)
	handlers.ListAllContactsUpdated(app, templates, db)
	handlers.Logout(app, templates)

	// Start the server
	address := ":" + port
	log.Printf("Starting server on %s %s", host, address)
	if err := app.Listen(address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
