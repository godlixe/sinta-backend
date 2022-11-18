package dto

type DetailTransaksiCreateDTO struct {
	ProdukID    uint64 `json:"produk_id" binding:"required"`
	TransaksiID uint64 `json:"transaksi_id"`
	TokoID      uint64 `json:"toko_id"`
	Jumlah      uint64 `json:"jumlah" binding:"required"`
	Harga       uint64 `json:"harga" binding:"required"`
}

type DetailTransaksiUpdateDTO struct {
	ID          uint64 `json:"id"`
	ProdukID    uint64 `json:"produk_id"`
	TransaksiID uint64 `json:"transaksi_id"`
	TokoID      uint64 `json:"toko_id"`
	Jumlah      uint64 `json:"jumlah"`
	Harga       uint64 `json:"harga"`
}
