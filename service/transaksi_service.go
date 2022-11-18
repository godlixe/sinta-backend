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
	stokRepository      repository.StokRepository
}

func NewTransaksiService(tr repository.TransaksiRepository, sr repository.StokRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: tr,
		stokRepository:      sr,
	}
}

func (s *transaksiService) CreateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDTO, tokoID uint64) (entity.Transaksi, error) {

	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(&transaksiDTO))
	if err != nil {
		return transaksi, err
	}

	for idx := range transaksi.DetailTransaksi {
		transaksi.DetailTransaksi[idx].TokoID = tokoID
	}

	for idx := range transaksi.DetailTransaksi {
		err := s.stokRepository.DecreaseStok(ctx, tokoID, transaksi.DetailTransaksi[idx].ProdukID, transaksi.DetailTransaksi[idx].Jumlah)
		if err != nil {
			return entity.Transaksi{}, err
		}
	}

	return s.transaksiRepository.CreateTransaksi(ctx, transaksi)
}

func (s *transaksiService) GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error) {
	return s.transaksiRepository.GetAllTransaksiByTokoID(ctx, tokoID)
}
