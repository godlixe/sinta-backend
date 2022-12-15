package repository

import (
	"context"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type KaryawanRepository interface {
	CreateKaryawan(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error)
	GetAllKaryawan(ctx context.Context) ([]entity.Karyawan, error)
	UpdateKaryawan(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error)
	DeleteKaryawan(ctx context.Context, karyawanID uint64) error
}

type karyawanConnection struct {
	connection *gorm.DB
}

func NewKaryawanRepository(db *gorm.DB) KaryawanRepository {
	return &karyawanConnection{
		connection: db,
	}
}

func (db *karyawanConnection) CreateKaryawan(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error) {
	tx := db.connection.Create(&karyawan)
	if tx.Error != nil {
		return entity.Karyawan{}, tx.Error
	}

	return karyawan, nil
}

func (db *karyawanConnection) GetAllKaryawan(ctx context.Context) ([]entity.Karyawan, error) {
	var daftarKaryawan []entity.Karyawan

	tx := db.connection.Find(&daftarKaryawan)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarKaryawan, nil
}

func (db *karyawanConnection) UpdateKaryawan(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error) {
	tx := db.connection.Updates(&karyawan)
	if tx.Error != nil {
		return entity.Karyawan{}, tx.Error
	}

	return karyawan, nil
}

func (db *karyawanConnection) DeleteKaryawan(ctx context.Context, karyawanID uint64) error {
	tx := db.connection.Where(("id = ?"), karyawanID).Delete(&entity.Karyawan{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
