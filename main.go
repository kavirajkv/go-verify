package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-verify/routes"
)

func main() {	
	app:=fiber.New()

	// Setup routes
	routes.Routes(app)
	
	
	fmt.Println("Starting server on :8000")
	app.Listen(":8000")

}
