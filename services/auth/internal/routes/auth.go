package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/handlers"
)

type Routes interface {
	Install(app *fiber.App)
}

type authRoutes struct {
	authHandler handlers.AuthHandler
}

func NewAuthRoute(authHandler handlers.AuthHandler) Routes {
	return &authRoutes{authHandler: authHandler}
}

func (r *authRoutes) Install(app *fiber.App) {
	prefix := "/auth"

	app.Post(prefix+"/register", r.authHandler.Register)
}
