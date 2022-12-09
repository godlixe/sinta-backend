package repository

import (
	"context"
	"sinta-backend/entity"

	"gorm.io/gorm"
)

type AjuanRepository interface {
	GetAllAjuan(ctx context.Context) ([]entity.Ajuan, error)
	GetAjuanByID(ctx context.Context, ajuanID uint64) (entity.Ajuan, error)
	CreateAjuan(ctx context.Context, ajuan entity.Ajuan) (entity.Ajuan, error)
	AcceptAjuan(ctx context.Context, ajuanID uint64) error
	DeclineAjuan(ctx context.Context, ajuanID uint64) error
}

type ajuanConnection struct {
	connection *gorm.DB
}

func NewAjuanRepository(db *gorm.DB) AjuanRepository {
	return &ajuanConnection{
		connection: db,
	}
}

func (db *ajuanConnection) GetAllAjuan(ctx context.Context) ([]entity.Ajuan, error) {
	var daftarAjuan []entity.Ajuan
	tx := db.connection.Find(&daftarAjuan)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarAjuan, nil
}

func (db *ajuanConnection) GetAjuanByID(ctx context.Context, ajuanID uint64) (entity.Ajuan, error) {
	var ajuan entity.Ajuan
	tx := db.connection.Preload("DetailAjuan").Find(&ajuan)
	if tx.Error != nil {
		return entity.Ajuan{}, tx.Error
	}

	return ajuan, nil
}

func (db *ajuanConnection) CreateAjuan(ctx context.Context, ajuan entity.Ajuan) (entity.Ajuan, error) {
	tx := db.connection.Where(("status = ?"), "false").Create(&ajuan)
	if tx.Error != nil {
		return entity.Ajuan{}, tx.Error
	}

	return ajuan, nil
}

func (db *ajuanConnection) AcceptAjuan(ctx context.Context, ajuanID uint64) error {
	tx := db.connection.Model(&entity.Ajuan{}).Where(("id = ?"), ajuanID).UpdateColumn("status", true)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *ajuanConnection) DeclineAjuan(ctx context.Context, ajuanID uint64) error {
	tx := db.connection.Model(&entity.Ajuan{}).Where(("id = ?"), ajuanID).UpdateColumn("status", false)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
