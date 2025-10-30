package handlers

import (
	"log"
	"main/components"

	"github.com/gofiber/fiber/v2"
)

func Services(title, iconPath string, app *fiber.App) {
	// Handle GET request - serve the services page
	app.Get("/services", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		component := components.Services(title, iconPath)
		if err := component.Render(c.Context(), c); err != nil {
			log.Printf("Error executing template: %v", err)
			return c.SendString(`
					<div id="result" 
						 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
						 hx-swap-oob="true">
					 ❌ Failed to render services page. Please try again.
					</div>
				`)
		}

		return nil
	})
}
