package main

import (
	"fmt"
	"go-verify/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func loadenv() {
	if os.Getenv("APP_ENV")!= "production" {
		fmt.Println("Loading .env file")
		err:=godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
		} else {
			fmt.Println("Loaded .env file")
		}
	}else{
		fmt.Println("Production environment detected, skipping .env loading")
	}
}
	
func main() {	
	loadenv()

	app:=fiber.New()

	routes.Routes(app)

	fmt.Println("Starting server on :8000")
	app.Listen(":8000")

}
