package controllers

import (
	"context"
	"net/http"
	"rest-project/app/models"
	"rest-project/app/models/request"
	"rest-project/app/services"
	"rest-project/app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(auth services.AuthService) *AuthController {
	return &AuthController{AuthService: auth}
}

func (ac *AuthController) RegisterUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req request.ReqRegister
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(models.ResError("Unprocessable Entity", map[string]interface{}{}))
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(models.ResError("Validation failed", fails))
	}
	// Hash password sebelum menyimpan ke database
	hashedPassword, errHash := utils.HashPassword(req.Password)
	if errHash != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResError("failed to hash password", map[string]interface{}{}))
	}
	req.Password = hashedPassword

	err := ac.AuthService.Register(c, &req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResError(err.Error(), map[string]interface{}{}))
	}
	return ctx.Status(http.StatusCreated).JSON(models.ResSuccess(map[string]interface{}{}))
}
