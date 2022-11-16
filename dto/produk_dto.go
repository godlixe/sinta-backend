package dto

type ProdukCreateDTO struct {
	Nama  string `json:"nama"`
	Harga uint   `json:"harga"`
}

type ProdukUpdateDTO struct {
	ID    uint64 `json:"id"`
	Nama  string `json:"nama"`
	Harga uint   `json:"harga"`
}
