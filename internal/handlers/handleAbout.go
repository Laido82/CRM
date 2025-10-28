package handlers

import (
	"log"
	"main/components"

	"github.com/gofiber/fiber/v2"
)

func About(app *fiber.App) {
	app.Get("/about", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		component := components.About()

		if err := component.Render(c.Context(), c); err != nil {
			log.Printf("Error rendering component: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	})

}
