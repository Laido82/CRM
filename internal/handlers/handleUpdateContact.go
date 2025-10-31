package handlers

import (
	"database/sql"
	"log"
	"main/internal/controllers"
	"main/internal/validators"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UpdateContact(title, iconPath string, app *fiber.App, db *sql.DB) {

	// Handle POST request - process form submission
	app.Post("/updateContact", func(c *fiber.Ctx) error {
		// Get id value from table row
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

		// Parse id and age

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Invalid ID format
					</div>
				`)
		}

		// Get other form values from addContact form
		firstName := strings.ToUpper(c.FormValue("firstName"))
		lastName := strings.ToUpper(c.FormValue("lastName"))
		email := c.FormValue("email")
		ageStr := c.FormValue("age")

		age, err := strconv.ParseUint(ageStr, 10, 64)
		if ageStr != "" {
			if err != nil {
				return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Invalid age format
					</div>
				`)
			}
		}
		// Validate email format if provided
		if email != "" {
			if !validators.IsValidEmail(email) {
				return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Invalid email format
					</div>
				`)
			}
		}
		// Validate age range if provided
		if ageStr != "" {
			if !validators.IsValidAge(uint64(age)) {
				return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Age must be between 18 and 120
					</div>
				`)
			}
		}

		// Check if contact exists
		contact, _ := controllers.GetContactByID(db, id)
		if contact.ID == 0 {
			return c.SendString(`
				<div id="result" 
					 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
					 hx-swap-oob="true">
				 ❌ Contact not found
				</div>
			`)
		}

		if firstName == "" {
			firstName = contact.FirstName
		}
		if lastName == "" {
			lastName = contact.LastName
		}
		if email == "" {
			email = contact.Email
		}

		if age == 0 {
			age = uint64(contact.Age)
		}

		// Update contact using controller
		err = controllers.UpdateContact(db, id, firstName, lastName, email, age)
		if err != nil {
			log.Printf("Error updating contact: %v %v ", contact, err)
			return c.SendString(`
		<div id="result" 
			 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
			 hx-swap-oob="true">
		 ❌ Error updating contact. Please try again.
		</div>
		`)
		}

		// Return success response for HTMX
		return c.SendString(`
		<div id="result" 
		     class="fixed bottom-5 right-5 bg-blue-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
		     hx-swap-oob="true">
		 🔄 Contact updated successfully!
		</div>
		`)
	})
}
