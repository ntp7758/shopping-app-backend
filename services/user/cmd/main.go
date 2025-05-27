package main

import (
	"net"

	"github.com/ntp7758/shopping-app-backend/libs/databases"
	g "github.com/ntp7758/shopping-app-backend/services/user/internal/grpc"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	dbClient, err := databases.NewMongoDBConnection()
	if err != nil {
		panic(err)
	}

	certFile := "path_to/server.crt"
	keyFile := "path_to/server.pem"
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
