package grpc

import "context"

type authServer struct {
}

func NewAuthServer() AuthServer {
	return &authServer{}
}

func (r *authServer) mustEmbedUnimplementedAuthServer()

func (r *authServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {

	res := RegisterResponse{
		Message: req.AuthId,
	}

	return &res, nil
}
