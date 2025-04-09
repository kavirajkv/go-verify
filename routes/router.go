package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the application
func Routes(app *fiber.App) {
	// Define your routes here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Server up and running"})
	})



}