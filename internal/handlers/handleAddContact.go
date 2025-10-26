package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"main/internal/controllers"
	"main/internal/validators"
	"strings"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddContact(app *fiber.App, templates *template.Template, db *sql.DB) {
	// Handle GET request - serve the form
	app.Get("/addContact", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")

		if err := templates.ExecuteTemplate(c, "contactsForm.html", nil); err != nil {
			log.Printf("Error executing template: %v", err)
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Failed to load page. Please try again.
					</div>
				`)
		}

		return nil
	})

	// Handle POST request - process form submission
	app.Post("/addContact", func(c *fiber.Ctx) error {
		// Get form values
		firstName := strings.ToUpper(c.FormValue("firstName"))
		lastName := strings.ToUpper(c.FormValue("lastName"))
		email := c.FormValue("email")
		ageStr := c.FormValue("age")

		// Validate input
		if firstName == "" || lastName == "" || email == "" || ageStr == "" {
			return c.SendString(`
							<div id="result" 
								 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
								 hx-swap-oob="true">
							 ❌ All fields are required
							</div>
							 `)
		}

		// Parse age
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Invalid age format
					</div>
					 `)
		}

		//Validate email format
		if !validators.IsValidEmail(email) {
			return c.SendString(`
		<div id="result" 
			 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
			 hx-swap-oob="true">
		 ❌ Invalid email format
		</div>
		 `)
		}

		//Validate age range
		if !validators.IsValidAge(uint64(age)) {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Age must be between 18 and 120
					</div>
				`)
		}

		// Add contact using controller
		err = controllers.InsertContact(db, firstName, lastName, email, uint64(age))
		if err != nil {
			log.Printf("Error adding contact: %v", err)
			return c.SendString(`
		<div id="result" 
			 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
			 hx-swap-oob="true">
		 ❌ Failed to add contact. Please try again.
		</div>
		 `)
		}

		// Return success response for HTMX
		return c.SendString(`
		<div id="result" 
		     class="fixed bottom-5 right-5 bg-green-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
		     hx-swap-oob="true">
		 ✅ Contact added successfully!
		</div>
		 `)
	})
}
