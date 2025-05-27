package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/libs/databases"
	g "github.com/ntp7758/shopping-app-backend/services/auth/internal/grpc"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/handlers"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/routes"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	dbClient, err := databases.NewMongoDBConnection()
	if err != nil {
		panic(err)
	}

	grpcServerHost := ""
	certFile := "path_to/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		panic(err)
	}

	cc, err := grpc.NewClient(grpcServerHost, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	grpcAuthClientService := g.NewAuthClientService(g.NewAuthClient(cc))

	authRepo, err := repository.NewAuthRepository(dbClient)
	if err != nil {
		panic(err)
	}

	authServ := services.NewAuthService(authRepo, grpcAuthClientService)

	authHand := handlers.NewAuthHandler(authServ)

	authRoute := routes.NewAuthRoute(authHand)

	app := fiber.New()

	authRoute.Install(app)

	app.Listen(":8080")
}
