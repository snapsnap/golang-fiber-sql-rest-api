package routes

import (
	"rest-project/app/controllers"
	"rest-project/app/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authService services.AuthService) {
	authController := controllers.NewAuthController(authService)

	app.Post("/register", authController.RegisterUser)
	// app.Post("/login", authController.Login)
}
