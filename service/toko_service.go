package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type TokoService interface {
	CreateToko(ctx context.Context, tokoDTO dto.TokoCreateDTO) (entity.Toko, error)
	GetTokoByUsername(ctx context.Context, username string) (entity.Toko, error)
	GetAllToko(ctx context.Context) ([]entity.Toko, error)
	UpdateToko(ctx context.Context, tokoDTO dto.TokoUpdateDTO) (entity.Toko, error)
	DeleteToko(ctx context.Context, tokoID uint64) error
}

type tokoService struct {
	tokoRepository repository.TokoRepository
}

func NewTokoService(tr repository.TokoRepository) TokoService {
	return &tokoService{
		tokoRepository: tr,
	}
}

func (s *tokoService) CreateToko(ctx context.Context, tokoDTO dto.TokoCreateDTO) (entity.Toko, error) {
	createdToko := entity.Toko{}
	err := smapping.FillStruct(&createdToko, smapping.MapFields(&tokoDTO))
	if err != nil {
		return createdToko, err
	}

	res, err := s.tokoRepository.CreateToko(ctx, createdToko)
	if err != nil {
		return createdToko, err
	}
	return res, nil
}

func (s *tokoService) GetTokoByUsername(ctx context.Context, username string) (entity.Toko, error) {
	return s.tokoRepository.GetTokoByUsername(ctx, username)
}

func (s *tokoService) GetAllToko(ctx context.Context) ([]entity.Toko, error) {
	return s.tokoRepository.GetAllToko(ctx)
}

func (s *tokoService) UpdateToko(ctx context.Context, tokoDTO dto.TokoUpdateDTO) (entity.Toko, error) {
	updatedToko := entity.Toko{}
	err := smapping.FillStruct(&updatedToko, smapping.MapFields(&tokoDTO))
	if err != nil {
		return updatedToko, err
	}

	res, err := s.tokoRepository.UpdateToko(ctx, updatedToko)
	if err != nil {
		return updatedToko, err
	}
	return res, nil
}

func (s *tokoService) DeleteToko(ctx context.Context, tokoID uint64) error {
	return s.tokoRepository.DeleteToko(ctx, tokoID)
}
