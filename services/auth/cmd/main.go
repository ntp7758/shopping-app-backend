package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ntp7758/shopping-app-backend/libs/databases"
	g "github.com/ntp7758/shopping-app-backend/services/auth/internal/grpc"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/handlers"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/routes"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/services"
	"github.com/ntp7758/shopping-app-backend/services/auth/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	utils.LoadConfig()
	dbURI := viper.GetString("DATABASE_URL")

	dbClient, err := databases.NewMongoDBConnection(dbURI)
	if err != nil {
		panic(err)
	}
	defer dbClient.DC()

	dbName := viper.GetString("DATABASE_NAME")
	err = dbClient.SetDB(dbName)
	if err != nil {
		panic(err)
	}

	grpcServerHost := viper.GetString("USER_GRPC_HOST")
	certFile := viper.GetString("CA_CERT_PATH")
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
	app.Use(cors.New())
	app.Use(logger.New())

	authRoute.Install(app)

	app.Listen(":8080")
}
