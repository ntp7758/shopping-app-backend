package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/domain"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/errs"
	g "github.com/ntp7758/shopping-app-backend/services/auth/internal/grpc"
	"github.com/ntp7758/shopping-app-backend/services/auth/internal/repository"
	"github.com/ntp7758/shopping-app-backend/services/auth/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Register(req domain.Register) error
}

type authService struct {
	authRepo              repository.AuthRepository
	grpcAuthClientService g.AuthClientService
}

func NewAuthService(authRepo repository.AuthRepository, grpcAuthClientService g.AuthClientService) AuthService {
	return &authService{authRepo: authRepo, grpcAuthClientService: grpcAuthClientService}
}

func (s *authService) Register(req domain.Register) error {
	if req.Password != req.ConfirmPassword {
		return errs.AppError{
			Code:    fiber.StatusBadRequest,
			Message: "password and confirm do not match",
		}
	}

	_, err := s.authRepo.GetByUsername(req.Username)
	if err == nil {
		return errs.AppError{
			Code:    fiber.StatusBadRequest,
			Message: "username is already used",
		}
	}
	if err != mongo.ErrNoDocuments {
		return errs.AppError{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	pwdHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return errs.AppError{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	t := time.Now()
	auth := domain.Auth{
		CreatedAt: t,
		UpdatedAt: t,
		Username:  req.Username,
		Password:  pwdHash,
	}

	id, err := s.authRepo.Insert(auth)
	if err != nil {
		return errs.AppError{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	err = s.grpcAuthClientService.Auth(id)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			return errs.AppError{
				Code:    int(grpcErr.Code()),
				Message: grpcErr.Message(),
			}
		}
		return errs.AppError{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}
