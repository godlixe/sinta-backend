package dto

type ProdukCreateDTO struct {
	Nama  string `json:"nama" binding:"required"`
	Harga uint   `json:"harga" binding:"required"`
}

type ProdukUpdateDTO struct {
	ID    uint64 `json:"id" binding:"required"`
	Nama  string `json:"nama" binding:"required"`
	Harga uint   `json:"harga" binding:"required"`
}
