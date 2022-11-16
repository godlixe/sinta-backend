package entity

type DetailTransaksi struct {
	ID          uint64     `json:"id"`
	ProdukID    uint64     `json:"produk_id" gorm:"foreignKey"`
	Produk      *Produk    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"produk,omitempty"`
	TransaksiID uint64     `json:"ajuan_id" gorm:"foreignKey"`
	Transaksi   *Transaksi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaksi,omitempty"`
	TokoID      uint64     `json:"toko_id" gorm:"foreignKey"`
	Toko        *Toko      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	Jumlah      uint64     `json:"jumlah"`
	Harga       uint64     `json:"harga"`
}
