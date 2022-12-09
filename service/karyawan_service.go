package service

import (
	"context"
	"sinta-backend/dto"
	"sinta-backend/entity"
	"sinta-backend/repository"

	"github.com/mashingan/smapping"
)

type KaryawanService interface {
	CreateKaryawan(ctx context.Context, karyawanDTO dto.KaryawanCreateDTO) (entity.Karyawan, error)
	GetAllKaryawan(ctx context.Context) ([]entity.Karyawan, error)
	UpdateKaryawan(ctx context.Context, karyawanDTO dto.KaryawanUpdateDTO) (entity.Karyawan, error)
	DeleteKaryawan(ctx context.Context, karyawanID uint64) error
}

type karyawanService struct {
	karyawanRepository repository.KaryawanRepository
}

func NewKaryawanService(kr repository.KaryawanRepository) KaryawanService {
	return &karyawanService{
		karyawanRepository: kr,
	}
}

func (s *karyawanService) CreateKaryawan(ctx context.Context, karyawanDTO dto.KaryawanCreateDTO) (entity.Karyawan, error) {
	karyawan := entity.Karyawan{}
	err := smapping.FillStruct(&karyawan, smapping.MapFields(&karyawanDTO))
	if err != nil {
		return karyawan, err
	}

	return s.karyawanRepository.CreateKaryawan(ctx, karyawan)
}

func (s *karyawanService) GetAllKaryawan(ctx context.Context) ([]entity.Karyawan, error) {
	return s.karyawanRepository.GetAllKaryawan(ctx)
}

func (s *karyawanService) UpdateKaryawan(ctx context.Context, karyawanDTO dto.KaryawanUpdateDTO) (entity.Karyawan, error) {
	karyawan := entity.Karyawan{}
	err := smapping.FillStruct(&karyawan, smapping.MapFields(&karyawanDTO))
	if err != nil {
		return karyawan, err
	}

	return s.karyawanRepository.UpdateKaryawan(ctx, karyawan)
}

func (s *karyawanService) DeleteKaryawan(ctx context.Context, karyawanID uint64) error {
	return s.karyawanRepository.DeleteKaryawan(ctx, karyawanID)
}
