package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type StokService interface {
	GetStokByTokoID(ctx context.Context, tokoID uint64) (entity.StokBatch, error)
	GetProdukStokByTokoID(ctx context.Context, tokoID uint64) ([]entity.StokToko, error)
	InsertStok(ctx context.Context, stokDTO dto.StokBatchCreateDTO, tokoID uint64) (entity.StokBatch, error)
	UpdateStok(ctx context.Context, stokDTO dto.StokBatchUpdateDTO, tokoID uint64) (entity.StokBatch, error)
	IncreaseStok(ctx context.Context, tokoID uint64, produkID uint64, amount uint64) error
	DecreaseStok(ctx context.Context, tokoID uint64, produkID uint64, sold uint64) error
}

type stokService struct {
	stokRepository repository.StokRepository
}

func NewStokService(sr repository.StokRepository) StokService {
	return &stokService{
		stokRepository: sr,
	}
}

func (s *stokService) GetStokByTokoID(ctx context.Context, tokoID uint64) (entity.StokBatch, error) {
	return s.stokRepository.GetStokByTokoID(ctx, tokoID)
}

func (s *stokService) GetProdukStokByTokoID(ctx context.Context, tokoID uint64) ([]entity.StokToko, error) {
	return s.stokRepository.GetProdukStokByTokoID(ctx, tokoID)
}

func (s *stokService) InsertStok(ctx context.Context, stokDTO dto.StokBatchCreateDTO, tokoID uint64) (entity.StokBatch, error) {
	stok := entity.StokBatch{}
	err := smapping.FillStruct(&stok, smapping.MapFields(&stokDTO))
	if err != nil {
		return stok, err
	}

	for idx := range stok.DaftarStok {
		stok.DaftarStok[idx].TokoID = tokoID
	}

	return s.stokRepository.InsertStok(ctx, stok)
}

func (s *stokService) UpdateStok(ctx context.Context, stokDTO dto.StokBatchUpdateDTO, tokoID uint64) (entity.StokBatch, error) {
	stok := entity.StokBatch{}
	err := smapping.FillStruct(&stok, smapping.MapFields(&stokDTO))
	if err != nil {
		return stok, err
	}

	for idx := range stok.DaftarStok {
		stok.DaftarStok[idx].TokoID = tokoID
	}

	return s.stokRepository.UpdateStok(ctx, stok)
}

func (s *stokService) IncreaseStok(ctx context.Context, tokoID uint64, produkID uint64, amount uint64) error {
	return s.stokRepository.IncreaseStok(ctx, tokoID, produkID, amount)
}

func (s *stokService) DecreaseStok(ctx context.Context, tokoID uint64, produkID uint64, sold uint64) error {
	return s.stokRepository.DecreaseStok(ctx, tokoID, produkID, sold)
}
