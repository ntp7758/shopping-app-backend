package grpc

import (
	"context"

	"github.com/ntp7758/shopping-app-backend/services/user/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	userService services.UserService
}

func NewAuthServer(userService services.UserService) AuthServer {
	return &authServer{userService: userService}
}

func (r *authServer) mustEmbedUnimplementedAuthServer() {}

func (r *authServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {

	if req.AuthId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "require auth id")
	}

	err := r.userService.Register(req.AuthId)
	if err != nil {
		return nil, err
	}

	res := RegisterResponse{
		Message: "success",
	}

	return &res, nil
}
