package service

import (
	// "echo-server/internal/model"
	"echo-server/internal/repository"
)

type UserService interface {
	FetchUsers() error
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) FetchUsers() error {
	// return s.userRepo.Create(userPayload)
	return nil
}
