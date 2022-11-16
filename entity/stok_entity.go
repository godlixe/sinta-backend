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
