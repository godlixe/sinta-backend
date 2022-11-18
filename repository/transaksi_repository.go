package repository

import (
	"context"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
	GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error)
}

type transaksiConnection struct {
	connection *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiConnection{
		connection: db,
	}
}

func (db *transaksiConnection) CreateTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error) {
	tx := db.connection.Create(&transaksi)
	if tx.Error != nil {
		return entity.Transaksi{}, tx.Error
	}

	return transaksi, nil
}

func (db *transaksiConnection) GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error) {
	var daftarTransaksi []entity.Transaksi
	tx := db.connection.Where(("toko_id"), tokoID).Find(&daftarTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarTransaksi, nil
}
