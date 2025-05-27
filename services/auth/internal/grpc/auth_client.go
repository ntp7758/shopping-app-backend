package grpc

import (
	"context"
	"log"
)

type AuthClientService interface {
	Auth(string) error
}

type authClientService struct {
	authClient AuthClient
}

func NewAuthClientService(authClient AuthClient) AuthClientService {
	return &authClientService{authClient: authClient}
}

func (c *authClientService) Auth(authId string) error {
	req := RegisterRequest{
		AuthId: authId,
	}

	res, err := c.authClient.Register(context.Background(), &req)
	if err != nil {
		return err
	}

	log.Println(res)

	return nil
}
