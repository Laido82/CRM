package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"main/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ListAllContactsUpdated(app *fiber.App, templates *template.Template, db *sql.DB) {
	app.Get("/listAllContactsUpdated", func(c *fiber.Ctx) error {
		// Set content type to HTML
		c.Set("Content-Type", "text/html")

		// Get contacts data
		contacts, err := controllers.GetAllContacts(db)
		if err != nil {
			log.Printf("Error getting contacts: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		// Execute template with contacts data
		if err := templates.ExecuteTemplate(c, "listAllContactsUpdated.html", contacts); err != nil {
			log.Printf("Error executing template: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	})
}
