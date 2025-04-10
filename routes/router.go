package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-verify/api"
)


func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Server up and running"})
	})

	app.Post("/send-otp", api.SendOTP)
	app.Post("/verify-otp", api.VerifyOTP)

}