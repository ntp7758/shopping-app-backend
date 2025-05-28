package main

import (
	"net"

	"github.com/ntp7758/shopping-app-backend/libs/databases"
	g "github.com/ntp7758/shopping-app-backend/services/user/internal/grpc"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/services"
	"github.com/ntp7758/shopping-app-backend/services/user/utils"
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

	certFile := viper.GetString("SERVER_CERT_PATH")
	keyFile := viper.GetString("SERVER_PEM_PATH")
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	userRepo, err := repository.NewUserRepository(dbClient)
	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(userRepo)

	grpcUserService := g.NewAuthServer(userService)

	g.RegisterAuthServer(grpcServer, grpcUserService)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
