package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mymodules/database"
	"github.com/mymodules/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
