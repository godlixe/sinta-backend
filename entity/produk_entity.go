package entity

type Produk struct {
	ID    uint64 `json:"id"`
	Nama  string `json:"nama"`
	Harga uint   `json:"harga"`
	BaseModel
}
