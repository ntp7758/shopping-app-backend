package services

import (
	"time"

	"github.com/ntp7758/shopping-app-backend/services/user/internal/domain"
	"github.com/ntp7758/shopping-app-backend/services/user/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Register(authId string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(authId string) error {

	_, err := s.userRepo.GetByAuthId(authId)
	if err == nil {
		return status.Errorf(codes.AlreadyExists, "auth id is already exists")
	}
	if err != mongo.ErrNoDocuments {
		return status.Errorf(codes.Internal, err.Error())
	}

	t := time.Now()
	user := domain.User{
		CreatedAt: t,
		UpdatedAt: t,
		AuthId:    authId,
	}

	_, err = s.userRepo.Insert(user)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
