package repository

import (
	"context"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type StokRepository interface {
	GetStokByTokoID(ctx context.Context, tokoID uint64) (entity.StokBatch, error)
	InsertStok(ctx context.Context, stok entity.StokBatch) (entity.StokBatch, error)
	UpdateStok(ctx context.Context, stok entity.StokBatch) (entity.StokBatch, error)
	IncreaseStok(ctx context.Context, tokoID uint64, produkID uint64, amount uint64) error
	DecreaseStok(ctx context.Context, tokoID uint64, produkID uint64, sold uint64) error
}

type stokConnection struct {
	connection *gorm.DB
}

func NewStokRepository(db *gorm.DB) StokRepository {
	return &stokConnection{
		connection: db,
	}
}

func (db *stokConnection) GetStokByTokoID(ctx context.Context, tokoID uint64) (entity.StokBatch, error) {
	var stok entity.StokBatch
	tx := db.connection.Where(("toko_id = ?"), tokoID).Preload("product").Find(&stok.DaftarStok)
	if tx.Error != nil {
		return entity.StokBatch{}, tx.Error
	}

	return stok, nil
}

func (db *stokConnection) InsertStok(ctx context.Context, stok entity.StokBatch) (entity.StokBatch, error) {
	tx := db.connection.Create(&stok.DaftarStok)
	if tx.Error != nil {
		return entity.StokBatch{}, tx.Error
	}

	return stok, nil
}

func (db *stokConnection) UpdateStok(ctx context.Context, stok entity.StokBatch) (entity.StokBatch, error) {
	tx := db.connection.Save(&stok.DaftarStok)
	if tx.Error != nil {
		return entity.StokBatch{}, tx.Error
	}

	return stok, nil
}

func (db *stokConnection) IncreaseStok(ctx context.Context, tokoID uint64, produkID uint64, amount uint64) error {
	tx := db.connection.Model(&entity.Stok{}).
		Where(("toko_id = ?"), tokoID).
		Where(("produk_id = ?"), produkID).
		UpdateColumn("jumlah", gorm.Expr("jumlah + ?", amount))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *stokConnection) DecreaseStok(ctx context.Context, tokoID uint64, produkID uint64, sold uint64) error {
	tx := db.connection.Model(&entity.Stok{}).
		Where(("toko_id = ?"), tokoID).
		Where(("produk_id = ?"), produkID).
		UpdateColumn("jumlah", gorm.Expr("jumlah - ?", sold))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
