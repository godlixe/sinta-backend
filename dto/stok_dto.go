package dto

type StokCreateDTO struct {
	TokoID   uint64 `json:"toko_id"`
	ProdukID uint64 `json:"produk_id"`
	Jumlah   uint64
}

type StokUpdateDTO struct {
	ID       uint64 `json:"id"`
	TokoID   uint64 `json:"toko_id"`
	ProdukID uint64 `json:"produk_id"`
	Jumlah   uint64
}

type StokBatchCreateDTO struct {
	DaftarStok []StokCreateDTO `json:"daftar_stok"`
}

type StokBatchUpdateDTO struct {
	DaftarStok []StokUpdateDTO `json:"daftar_stok"`
}
