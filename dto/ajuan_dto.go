package dto

type AjuanCreateDTO struct {
	Status bool   `json:"status"`
	TokoID uint64 `json:"toko_id"`
}

type AjuanUpdateDTO struct {
	ID     uint64 `json:"id"`
	Status bool   `json:"status"`
	TokoID uint64 `json:"toko_id"`
}
