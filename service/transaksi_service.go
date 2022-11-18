package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO, tokoID uint64) (entity.Transaksi, error)
	GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error)
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(tr repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: tr,
	}
}

func (s *transaksiService) CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO, tokoID uint64) (entity.Transaksi, error) {

	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(&transaksiDTO))
	if err != nil {
		return transaksi, err
	}

	for idx, _ := range transaksi.DetailTransaksi {
		transaksi.DetailTransaksi[idx].TokoID = tokoID
	}
	return s.transaksiRepository.CreateTransaksi(ctx, transaksi)
}

func (s *transaksiService) GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error) {
	return s.transaksiRepository.GetAllTransaksiByTokoID(ctx, tokoID)
}
