package dto

type AjuanCreateDTO struct {
	Status      bool                   `json:"status"`
	TokoID      uint64                 `json:"toko_id"`
	DetailAjuan []DetailAjuanCreateDTO `json:"detail_ajuan"`
}

type AjuanUpdateDTO struct {
	ID          uint64                 `json:"id"`
	Status      bool                   `json:"status"`
	TokoID      uint64                 `json:"toko_id"`
	DetailAjuan []DetailAjuanUpdateDTO `json:"detail_ajuan"`
}
