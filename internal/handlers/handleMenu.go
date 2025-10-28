package handlers

import (
	"main/components"

	"github.com/gofiber/fiber/v2"
)

// Menu handler to render the menu page
func Menu(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		component := components.Menu()
		if err := component.Render(c.Context(), c); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
				<div id="result" 
					 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
					 hx-swap-oob="true">
				 ❌ Error rendering menu
				</div>
			`)
		}
		return nil
	})
}
