package entity

type Stok struct {
	ID       uint64  `json:"id"`
	TokoID   uint64  `json:"toko_id" gorm:"foreignKey"`
	Toko     *Toko   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	ProdukID uint64  `json:"produk_id" gorm:"foreignKey"`
	Produk   *Produk `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"produk,omitempty"`
	Jumlah   uint64
	BaseModel
}

type StokBatch struct {
	DaftarStok []Stok `json:"daftar_stok"`
}

type StokToko struct {
	ID     uint64 `json:"id"`
	Nama   string `json:"nama"`
	Jumlah uint64 `json:"jumlah"`
	Harga  uint64 `json:"harga"`
}
