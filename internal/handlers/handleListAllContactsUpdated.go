package handlers

import (
	"database/sql"
	"log"
	"main/components"
	"main/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ListAllContactsUpdated(app *fiber.App,  db *sql.DB) {
	app.Get("/listAllContactsUpdated", func(c *fiber.Ctx) error {
		// Set content type to HTML
		c.Set("Content-Type", "text/html")

		// Get contacts data
		contacts, err := controllers.GetAllContacts(db)
		if err != nil {
			log.Printf("Error getting contacts: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}

		component := components.ListAllContactsUpdated(contacts)
		if err := component.Render(c.Context(), c); err != nil {
			log.Printf("Error rendering template: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	})
}
