package repository

import (
	"context"
	"errors"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	CreateProduk(ctx context.Context, produk entity.Produk) (entity.Produk, error)
	GetAllProduk(ctx context.Context) ([]entity.Produk, error)
	UpdateProduk(ctx context.Context, produk entity.Produk) (entity.Produk, error)
	DeleteProduk(ctx context.Context, produkID uint64) error
}

type produkConnection struct {
	connection *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkConnection{
		connection: db,
	}
}

func (db *produkConnection) CreateProduk(ctx context.Context, produk entity.Produk) (entity.Produk, error) {
	tx := db.connection.Create(&produk)
	if tx.Error != nil {
		return entity.Produk{}, tx.Error
	}

	return produk, nil
}

func (db *produkConnection) GetAllProduk(ctx context.Context) ([]entity.Produk, error) {
	var daftarProduk []entity.Produk

	tx := db.connection.Find(&daftarProduk)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(daftarProduk) <= 0 {
		return nil, errors.New("no products found")
	}

	return daftarProduk, nil
}

func (db *produkConnection) UpdateProduk(ctx context.Context, produk entity.Produk) (entity.Produk, error) {
	tx := db.connection.Save(&produk)
	if tx.Error != nil {
		return entity.Produk{}, tx.Error
	}

	return produk, nil
}

func (db *produkConnection) DeleteProduk(ctx context.Context, produkID uint64) error {
	tx := db.connection.Where(("id"), produkID).Delete(&entity.Produk{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
