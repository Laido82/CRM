package handlers

import (
	"database/sql"
	"log"
	"main/components"
	"main/internal/controllers"
	"main/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ListSearchedContacts(app *fiber.App, db *sql.DB) {
	// Handler GET request - serve the form
	app.Get("/listSearchedContacts", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		// Get contacts data
		contacts, err := controllers.GetAllContacts(db)
		if err != nil {
			log.Printf("Error getting contacts: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		component := components.ListSearchedContacts(contacts)
		if err := component.Render(c.Context(), c); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Error rendering contacts
					</div>
				`)
		}
		return nil
	})

	//Handle POST request - process form submission
	app.Post("/searchedContacts", func(c *fiber.Ctx) error {
		// Get form values
		searchTerm := strings.ToUpper(c.FormValue("searchTerm"))
		// If search term is empty, set it to wildcard to fetch all contacts
		if searchTerm == "" {
			searchTerm = "%"
		}

		// Query the database for contacts matching the search term
		var contacts []models.Contact
		contacts, err := controllers.SearchedContacts(db, searchTerm)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Error loading contacts
					</div>
				`)
		}
		// If no contacts found, inform the user
		if len(contacts) == 0 {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-yellow-500 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ⚠️ No contacts found matching the search criteria
					</div>
				`)
		}

		// Render the results
		c.Set("Content-Type", "text/html")
		component := components.ListSearchedContacts(contacts)
		if err := component.Render(c.Context(), c); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Error rendering contacts
					</div>
				`)
		}

		return nil
	})
}
