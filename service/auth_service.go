package service

import (
	"context"
	"sinta-backend/helpers"
	"sinta-backend/repository"
)

type AuthService interface {
	VerifyCredential(ctx context.Context, email string, password string) (bool, error)
	CheckEmailDuplicate(ctx context.Context, email string) (bool, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(ur repository.UserRepository) AuthService {
	return &authService{
		userRepository: ur,
	}
}

func (s *authService) VerifyCredential(ctx context.Context, email string, password string) (bool, error) {
	res, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	comparedPassword, err := helpers.ComparePassword(res.Password, []byte(password))
	if err != nil {
		return false, err
	}

	if res.Email == email && comparedPassword {
		return true, nil
	}

	return false, nil
}

func (s *authService) CheckEmailDuplicate(ctx context.Context, email string) (bool, error) {
	res, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if res.Email == "" {
		return false, nil
	}
	return true, nil
}
