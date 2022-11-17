package repository

import (
	"context"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type TokoRepository interface {
	CreateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error)
	GetTokoByUsername(ctx context.Context, username string) (entity.Toko, error)
	UpdateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error)
	DeleteToko(ctx context.Context, tokoID uint64) error
}

type tokoConnection struct {
	connection *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoConnection{
		connection: db,
	}
}

func (db *tokoConnection) CreateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error) {
	tx := db.connection.Create(&toko)
	if tx.Error != nil {
		return entity.Toko{}, tx.Error
	}

	return toko, nil
}

func (db *tokoConnection) GetTokoByUsername(ctx context.Context, username string) (entity.Toko, error) {
	var toko entity.Toko
	tx := db.connection.Where(("username = ?"), username).Take(&toko)
	if tx.Error != nil {
		return toko, tx.Error
	}
	return toko, nil
}

func (db *tokoConnection) UpdateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error) {
	tx := db.connection.Save(&toko)
	if tx.Error != nil {
		return entity.Toko{}, tx.Error
	}

	return toko, nil
}

func (db *tokoConnection) DeleteToko(ctx context.Context, tokoID uint64) error {
	tx := db.connection.Delete(&entity.Toko{}, tokoID)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
