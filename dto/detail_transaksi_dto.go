package dto

type DetailTransaksiCreateDTO struct {
	ProdukID    uint64 `json:"produk_id"`
	TransaksiID uint64 `json:"ajuan_id"`
	TokoID      uint64 `json:"toko_id"`
	Jumlah      uint64 `json:"jumlah"`
	Harga       uint64 `json:"harga"`
}

type DetailTransaksiUpdateDTO struct {
	ID          uint64 `json:"id"`
	ProdukID    uint64 `json:"produk_id"`
	TransaksiID uint64 `json:"ajuan_id"`
	TokoID      uint64 `json:"toko_id"`
	Jumlah      uint64 `json:"jumlah"`
	Harga       uint64 `json:"harga"`
}
