package repository

import (
	"context"
	"sinta-backend/entity"
	"time"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
	GetAllTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.Transaksi, error)
	GetHarianTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error)
	GetMingguanTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error)
	GetBulananTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error)
	GetHarianTotal(ctx context.Context) ([]entity.TransaksiTotal, error)
	GetMingguanTotal(ctx context.Context) ([]entity.TransaksiTotal, error)
	GetBulananTotal(ctx context.Context) ([]entity.TransaksiTotal, error)
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
	tx := db.connection.Where(("toko_id"), tokoID).Preload("DetailTransaksi").Find(&daftarTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarTransaksi, nil
}

func (db *transaksiConnection) GetHarianTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error) {
	var daftarTransaksi []entity.DetailTransaksi

	now := time.Now()
	currentYear, currentMonth, currDate := now.Date()
	currentLocation := now.Location()

	firstOfDay := time.Date(currentYear, currentMonth, currDate, 0, 0, 0, 0, currentLocation)
	lastOfDay := firstOfDay.AddDate(0, 0, 1)

	subQueryHarian := db.connection.Model(&entity.Transaksi{}).Select("id").Where(("toko_id = ?"), tokoID).Where("created_at >= ?", firstOfDay).Where("created_at <= ?", lastOfDay)
	tx := db.connection.Model(&entity.DetailTransaksi{}).Where("transaksi_id IN (?)", subQueryHarian).Preload("Produk").Find(&daftarTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarTransaksi, nil
}

func (db *transaksiConnection) GetMingguanTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error) {
	var daftarTransaksi []entity.DetailTransaksi
	currTime := time.Now()
	currTime = currTime.AddDate(0, 0, -7)
	subQueryMingguan := db.connection.Model(&entity.Transaksi{}).Select("id").Where(("toko_id = ?"), tokoID).Where("created_at >= ?", currTime)
	tx := db.connection.Model(&entity.DetailTransaksi{}).Where("transaksi_id IN (?)", subQueryMingguan).Preload("Produk").Find(&daftarTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarTransaksi, nil
}

func (db *transaksiConnection) GetBulananTransaksiByTokoID(ctx context.Context, tokoID uint64) ([]entity.DetailTransaksi, error) {
	var daftarTransaksi []entity.DetailTransaksi

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	subQueryBulanan := db.connection.Model(&entity.Transaksi{}).Select("id").Where(("toko_id = ?"), tokoID).Where("created_at >= ?", firstOfMonth).Where("created_at <= ?", lastOfMonth)
	tx := db.connection.Model(&entity.DetailTransaksi{}).Where("transaksi_id = (?)", subQueryBulanan).Preload("Produk").Find(&daftarTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return daftarTransaksi, nil
}

func (db *transaksiConnection) GetHarianTotal(ctx context.Context) ([]entity.TransaksiTotal, error) {
	var totalTransaksi []entity.TransaksiTotal
	now := time.Now()
	currentYear, currentMonth, currDate := now.Date()
	currentLocation := now.Location()

	firstOfDay := time.Date(currentYear, currentMonth, currDate, 0, 0, 0, 0, currentLocation)
	lastOfDay := firstOfDay.AddDate(0, 0, 1)

	subQueryHarian := db.connection.Model(&entity.Transaksi{}).Select("id").Where("created_at >= ?", firstOfDay).Where("created_at <= ?", lastOfDay)
	rows, err := db.connection.Model(&entity.DetailTransaksi{}).Select("toko_id, SUM(harga)").Having("transaksi_id = (?)", subQueryHarian).Group("toko_id").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		total := entity.TransaksiTotal{}
		if err := rows.Scan(&total.NamaToko, &total.Total); err != nil {
			return totalTransaksi, err
		}
		totalTransaksi = append(totalTransaksi, total)
	}
	return totalTransaksi, nil
}

func (db *transaksiConnection) GetMingguanTotal(ctx context.Context) ([]entity.TransaksiTotal, error) {
	var totalTransaksi []entity.TransaksiTotal
	currTime := time.Now()
	currTime = currTime.AddDate(0, 0, -7)

	subQueryMingguan := db.connection.Model(&entity.Transaksi{}).Select("id").Where("created_at >= ?", currTime)
	rows, err := db.connection.Model(&entity.DetailTransaksi{}).Select("toko_id, SUM(harga)").Having("transaksi_id = (?)", subQueryMingguan).Group("toko_id").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		total := entity.TransaksiTotal{}
		if err := rows.Scan(&total.NamaToko, &total.Total); err != nil {
			return totalTransaksi, err
		}
		totalTransaksi = append(totalTransaksi, total)
	}

	return totalTransaksi, nil
}

func (db *transaksiConnection) GetBulananTotal(ctx context.Context) ([]entity.TransaksiTotal, error) {
	var totalTransaksi []entity.TransaksiTotal

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	subQueryBulanan := db.connection.Model(&entity.Transaksi{}).Select("id").Where("created_at >= ?", firstOfMonth).Where("created_at <= ?", lastOfMonth)
	rows, err := db.connection.Model(&entity.DetailTransaksi{}).Select("toko_id, SUM(harga)").Having("transaksi_id = (?)", subQueryBulanan).Group("toko_id").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		total := entity.TransaksiTotal{}
		if err := rows.Scan(&total.NamaToko, &total.Total); err != nil {
			return totalTransaksi, err
		}
		totalTransaksi = append(totalTransaksi, total)
	}

	return totalTransaksi, nil
}
