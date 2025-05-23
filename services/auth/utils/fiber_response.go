package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/errs"
)

func FiberErrorResponse(c *fiber.Ctx, err error) error {

	var code int
	var msg string
	switch e := err.(type) {
	case errs.AppError:
		code = e.Code
		msg = e.Message
	case *fiber.Error:
		code = e.Code
		msg = e.Message
	case error:
		code = fiber.StatusInternalServerError
		msg = e.Error()
	}

	return c.Status(code).JSON(fiber.Map{
		"message": msg,
	})
}

func FiberSuccessResponse(c *fiber.Ctx, code int, data interface{}) error {
	return c.Status(code).JSON(data)
}
