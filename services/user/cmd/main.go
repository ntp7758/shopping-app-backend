package main

import (
	"net"

	g "github.com/ntp7758/shopping-app-backend/services/user/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

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

	g.RegisterAuthServer(grpcServer, g.NewAuthServer())

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
