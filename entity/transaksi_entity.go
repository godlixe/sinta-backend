package entity

type Transaksi struct {
	ID     uint64 `json:"id"`
	TokoID uint64 `json:"toko_id" gorm:"foreignKey"`
	Toko   *Toko  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	Total  uint64 `json:"total"`
	BaseModel
}
