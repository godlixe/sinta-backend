package dto

type TransaksiCreateDTO struct {
	TokoID          uint64                     `json:"toko_id"`
	DetailTransaksi []DetailTransaksiCreateDTO `json:"detail_transaksi" binding:"required"`
}

type TransaksiUpdateDTO struct {
	ID              uint64                     `json:"id" binding:"required"`
	TokoID          uint64                     `json:"toko_id"`
	DetailTransaksi []DetailTransaksiUpdateDTO `json:"detail_transaksi" binding:"required"`
}
