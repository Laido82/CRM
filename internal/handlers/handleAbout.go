package handlers

import (
	"log"
	"main/components"

	"github.com/gofiber/fiber/v2"
)

func About(title, iconPath string, app *fiber.App) {
	app.Get("/about", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		component := components.About(title, iconPath)

		if err := component.Render(c.Context(), c); err != nil {
			log.Printf("Error rendering component: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	})

}
