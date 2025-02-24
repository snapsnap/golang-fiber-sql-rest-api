package routes

import (
	"database/sql"
	"rest-project/app/repositories"
	"rest-project/app/services"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App, db *sql.DB) {
	userRepo := repositories.NewUserRepoImpl(db)
	authService := services.NewAuthServiceImpl(db, userRepo)

	AuthRoutes(app, authService)
}
