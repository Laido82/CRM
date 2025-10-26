package handlers

import (
	"github.com/gofiber/fiber/v2"
	"html/template"
)

// Menu handler to render the menu page
func Menu(app *fiber.App, templates *template.Template) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		if err := templates.ExecuteTemplate(c, "menu.html", nil); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
				<div id="result" 
					 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
					 hx-swap-oob="true">
				 ❌ Error loading menu
				</div>
			`)
		}
		return nil
	})
}
