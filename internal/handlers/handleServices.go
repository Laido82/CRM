package handlers

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandlerServices(app *fiber.App, templates *template.Template) {
	// Handle GET request - serve the services page
	app.Get("/services", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")

		if err := templates.ExecuteTemplate(c, "services.html", nil); err != nil {
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
}
