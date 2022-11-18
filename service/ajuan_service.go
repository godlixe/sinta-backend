package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type AjuanService interface {
	GetAllAjuan(ctx context.Context) ([]entity.Ajuan, error)
	GetAjuanByID(ctx context.Context, ajuanID uint64) (entity.Ajuan, error)
	CreateAjuan(ctx context.Context, tokoID uint64, ajuanDTO dto.AjuanCreateDTO) (entity.Ajuan, error)
	AcceptAjuan(ctx context.Context, ajuanID uint64) error
	DeclineAjuan(ctx context.Context, ajuanID uint64) error
}

type ajuanService struct {
	ajuanRepository repository.AjuanRepository
	stokRepository  repository.StokRepository
}

func NewAjuanService(ar repository.AjuanRepository, sr repository.StokRepository) AjuanService {
	return &ajuanService{
		ajuanRepository: ar,
		stokRepository:  sr,
	}
}

func (s *ajuanService) GetAllAjuan(ctx context.Context) ([]entity.Ajuan, error) {
	return s.ajuanRepository.GetAllAjuan(ctx)
}

func (s *ajuanService) GetAjuanByID(ctx context.Context, ajuanID uint64) (entity.Ajuan, error) {
	return s.ajuanRepository.GetAjuanByID(ctx, ajuanID)
}

func (s *ajuanService) CreateAjuan(ctx context.Context, tokoID uint64, ajuanDTO dto.AjuanCreateDTO) (entity.Ajuan, error) {
	ajuan := entity.Ajuan{}
	err := smapping.FillStruct(&ajuan, smapping.MapFields(&ajuanDTO))
	if err != nil {
		return ajuan, err
	}

	for idx := range ajuan.DetailAjuan {
		ajuan.DetailAjuan[idx].TokoID = tokoID
	}

	return s.ajuanRepository.CreateAjuan(ctx, ajuan)
}

func (s *ajuanService) AcceptAjuan(ctx context.Context, ajuanID uint64) error {
	res, err := s.ajuanRepository.GetAjuanByID(ctx, ajuanID)
	if err != nil {
		return err
	}

	for idx := range res.DetailAjuan {
		err := s.stokRepository.IncreaseStok(ctx, res.TokoID, res.DetailAjuan[idx].ProdukID, res.DetailAjuan[idx].Jumlah)
		if err != nil {
			return err
		}
	}

	return s.ajuanRepository.AcceptAjuan(ctx, ajuanID)
}

func (s *ajuanService) DeclineAjuan(ctx context.Context, ajuanID uint64) error {
	return s.ajuanRepository.DeclineAjuan(ctx, ajuanID)
}
