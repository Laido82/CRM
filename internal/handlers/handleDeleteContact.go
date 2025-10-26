package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"main/internal/controllers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteContact(app *fiber.App, templates *template.Template, db *sql.DB) {
	// Handle GET request - serve the form
	app.Get("/deleteContact", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")

		if err := templates.ExecuteTemplate(c, "contactsForm.html", nil); err != nil {
			log.Printf("Error executing template: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Error loading page
					</div>
				`)
		}

		return nil
	})

	// Handle POST request - process form submission
	app.Post("/deleteContact", func(c *fiber.Ctx) error {
		// Get form values
		idStr := c.FormValue("id")

		// Validate input
		if idStr == "" {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ ID is required
					</div>
				`)
		}

		// Parse id
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.SendString(`
				<div id="result" 
					 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
					 hx-swap-oob="true">
				 ❌ Invalid ID format
				</div>
			`)
		}
		contact, _ := controllers.GetContactByID(db, uint64(id))
		// Get contact before deleteing and Check if contact exists
		if _, err := controllers.GetContactByID(db, uint64(id)); err != nil || contact.ID == 0 {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Contact not found
					</div>
				`)
		}

		// Delete contact using controller
		if err = controllers.DeleteContact(db, contact.ID); err != nil {
			log.Printf("Error deleting contact: %v %v ", contact, err)
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Error deleting contact
					</div>
				`)
		}

		// Return success response for HTMX
		return c.SendString(`
		<div id="result" class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
		     hx-swap-oob="true">
		 🗑️ Contact deleted successfully!
		</div>
		`)
	})
}
