package handlers

import (
	"github.com/gofiber/fiber/v2"
	"html/template"
	"log"
)

func About(app *fiber.App, templates *template.Template) {
	app.Get("/about", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		if err := templates.ExecuteTemplate(c, "about.html", nil); err != nil {
			log.Printf("Error executing template: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	})
}
