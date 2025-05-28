package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/domain"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/errs"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/services"
	"github.com/ntp7758/shopping-app-backend/services/auth/utils"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
}

type authHandler struct {
	authServ services.AuthService
}

func NewAuthHandler(authServ services.AuthService) AuthHandler {
	return &authHandler{authServ: authServ}
}

func (h *authHandler) Register(c *fiber.Ctx) error {

	var req domain.Register
	err := c.BodyParser(&req)
	if err != nil {
		return utils.FiberErrorResponse(c, err)
	}

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" || strings.TrimSpace(req.ConfirmPassword) == "" {
		return utils.FiberErrorResponse(c, errs.AppError{
			Code:    fiber.StatusBadRequest,
			Message: "have input empty",
		})
	}

	err = h.authServ.Register(req)
	if err != nil {
		return utils.FiberErrorResponse(c, err)
	}

	return utils.FiberSuccessResponse(c, fiber.StatusCreated, map[string]interface{}{
		"message": "register success",
	})
}
