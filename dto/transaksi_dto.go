package dto

type TransaksiCreateDTO struct {
	TokoID uint64 `json:"toko_id"`
	Total  uint64 `json:"total"`
}

type TransaksiUpdateDTO struct {
	ID     uint64 `json:"id"`
	TokoID uint64 `json:"toko_id"`
	Total  uint64 `json:"total"`
}
