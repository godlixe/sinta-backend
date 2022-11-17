package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type ProdukService interface {
	CreateProduk(ctx context.Context, produkDTO dto.ProdukCreateDTO) (entity.Produk, error)
	GetAllProduk(ctx context.Context) ([]entity.Produk, error)
	UpdateProduk(ctx context.Context, produkDTO dto.ProdukUpdateDTO) (entity.Produk, error)
	DeleteProduk(ctx context.Context, produkID uint64) error
}

type produkService struct {
	produkRepository repository.ProdukRepository
}

func NewProdukService(pr repository.ProdukRepository) ProdukService {
	return &produkService{
		produkRepository: pr,
	}
}

func (s *produkService) CreateProduk(ctx context.Context, produkDTO dto.ProdukCreateDTO) (entity.Produk, error) {
	produk := entity.Produk{}
	err := smapping.FillStruct(&produk, smapping.MapFields(&produkDTO))
	if err != nil {
		return produk, err
	}

	return s.produkRepository.CreateProduk(ctx, produk)
}

func (s *produkService) GetAllProduk(ctx context.Context) ([]entity.Produk, error) {
	return s.produkRepository.GetAllProduk(ctx)
}

func (s *produkService) UpdateProduk(ctx context.Context, produkDTO dto.ProdukUpdateDTO) (entity.Produk, error) {
	produk := entity.Produk{}
	err := smapping.FillStruct(&produk, smapping.MapFields(&produkDTO))
	if err != nil {
		return produk, err
	}

	return s.produkRepository.UpdateProduk(ctx, produk)
}

func (s *produkService) DeleteProduk(ctx context.Context, produkID uint64) error {
	return s.produkRepository.DeleteProduk(ctx, produkID)
}
