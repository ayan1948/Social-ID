package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mymodules/database"
	"github.com/mymodules/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) //Making the frontend and the backend to run on different ports
	routes.Setup(app)
	app.Listen(":8000")
}
