package dto

type DetailAjuanCreateDTO struct {
	ProdukID uint64 `json:"produk_id"`
	AjuanID  uint64 `json:"ajuan_id"`
	TokoID   uint64 `json:"toko_id"`
	Jumlah   uint64 `json:"jumlah"`
}

type DetailAjuanUpdateDTO struct {
	ID       uint64 `json:"id"`
	ProdukID uint64 `json:"produk_id"`
	AjuanID  uint64 `json:"ajuan_id"`
	TokoID   uint64 `json:"toko_id"`
	Jumlah   uint64 `json:"jumlah"`
}
