package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mymodules/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.Profile)
	app.Get("/api/:id", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
