package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/libs/databases"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/handlers"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/routes"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/services"
)

func main() {
	dbClient, err := databases.NewMongoDBConnection()
	if err != nil {
		panic(err)
	}

	authRepo, err := repository.NewAuthRepository(dbClient)
	if err != nil {
		panic(err)
	}

	authServ := services.NewAuthService(authRepo)

	authHand := handlers.NewAuthHandler(authServ)

	authRoute := routes.NewAuthRoute(authHand)

	app := fiber.New()

	authRoute.Install(app)

	app.Listen(":8080")
}
