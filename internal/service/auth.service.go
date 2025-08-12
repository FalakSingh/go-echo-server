package service

import (
	appError "echo-server/internal/helper/errors"
	"echo-server/internal/model"
	"echo-server/internal/repository"
	"fmt"
)

type AuthService interface {
	CreateUser(payload *model.User) error
	LoginUser(payload *model.User) error
}

type authSvc struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authSvc{
		userRepo: userRepo,
	}
}

func (s *authSvc) CreateUser(payload *model.User) error {
	emailExists, err := s.userRepo.FindByEmail(payload.Email)
	if err != nil {
		return err
	}
	if emailExists != nil {
		return appError.NewBadRequest("Email already exists")
	}
	return s.userRepo.Create(payload)
}

func (s *authSvc) LoginUser(payload *model.User) error {
	fmt.Printf("payload: %+v", payload)
	user, err := s.userRepo.FindByEmail(payload.Email)
	fmt.Printf("user: %+v", user)
	if err != nil {
		return err
	}
	if user == nil {
		return appError.NewBadRequest("User not found")
	}
	*payload = *user
	return nil
}
