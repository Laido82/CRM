package handlers

import (
	"html/template"
	"main/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// HandleLogout renders the logout page and handles logout action
func Logout(app *fiber.App, templates *template.Template) {

	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Set("content-type", "text/html")
		if err := templates.ExecuteTemplate(c, "logout.html", nil); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(`
				<div id="result" 
					 class="fixed bottom-5 right-5 bg-red-600 text-white px-4 py-2 rounded-lg shadow-lg popup" 
					 hx-swap-oob="true">
					✗ Failed to load logout page
				</div>
			`)
		}

		// Shutdown the app gracefully

		go func() {
			controllers.MakeCopyDB()
			app.Shutdown()
		}()

		return c.SendString(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<link rel="icon" type="image/png" href="/static/icons/logout.png" />
			<title>Logging out...</title>
			<style>
				* {
					margin: 0;
					padding: 0;
					box-sizing: border-box;
				}
				body {
					background-color: #0f172b;
					min-height: 100vh;
					display: flex;
					align-items: center;
					justify-content: center;
					font-family: system-ui, -apple-system, sans-serif;
				}
				.container {
					background-color: white;
					border-radius: 1rem;
					box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
					padding: 2rem 3rem;
					text-align: center;
					animation: bounce 1s infinite;
				}
				.icon-wrapper {
					margin-bottom: 1rem;
				}
				.icon {
					width: 5rem;
					height: 5rem;
					margin: 0 auto;
					color: red;
				}
				h1 {
					font-size: 3rem;
					font-weight: bold;
					color: red;
					margin-bottom: 0.5rem;
				}
				h4 {
					color: green;
					font-size: 1.5rem;
					font-weight: normal;
				}
				@keyframes bounce {
					0%, 100% {
						transform: translateY(-5%);
						animation-timing-function: cubic-bezier(0.8, 0, 1, 1);
					}
					50% {
						transform: translateY(0);
						animation-timing-function: cubic-bezier(0, 0, 0.2, 1);
					}
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="icon-wrapper">
					<svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
					</svg>
				</div>
				<h1>Successfully Logged Out</h1>
				<h4>Thank you for using our service</h4>
			</div>
		</body>
		</html>
		`)
	})

}
