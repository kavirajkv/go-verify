package main

import (
	"fmt"
	"go-verify/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func loadenv() {
	err:=godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Loaded .env file")
	}
}
	
func main() {	
	loadenv()

	app:=fiber.New()

	routes.Routes(app)

	fmt.Println("Starting server on :8000")
	app.Listen(":8000")

}
