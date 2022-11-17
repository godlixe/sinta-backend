package service

import (
	"context"
	"sinta-backend/helpers"
	"sinta-backend/repository"
)

type AuthService interface {
	VerifyCredential(ctx context.Context, username string, password string) (bool, error)
	CheckUsernameDuplicate(ctx context.Context, username string) (bool, error)
}

type authService struct {
	tokoRepository repository.TokoRepository
}

func NewAuthService(tr repository.TokoRepository) AuthService {
	return &authService{
		tokoRepository: tr,
	}
}

func (s *authService) VerifyCredential(ctx context.Context, username string, password string) (bool, error) {
	res, err := s.tokoRepository.GetTokoByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	comparedPassword, err := helpers.ComparePassword(res.Password, []byte(password))
	if err != nil {
		return false, err
	}

	if res.Username == username && comparedPassword {
		return true, nil
	}

	return false, nil
}

func (s *authService) CheckUsernameDuplicate(ctx context.Context, username string) (bool, error) {
	res, err := s.tokoRepository.GetTokoByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	if res.Username == "" {
		return false, nil
	}
	return true, nil
}
