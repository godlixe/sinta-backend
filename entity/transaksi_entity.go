package entity

type Transaksi struct {
	ID              uint64            `json:"id"`
	TokoID          uint64            `json:"toko_id" gorm:"foreignKey"`
	Toko            *Toko             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	DetailTransaksi []DetailTransaksi `json:"detail_transaksi"`
	BaseModel
}

type TransaksiTotal struct {
	NamaToko string `json:"nama_toko"`
	Total    int    `json:"total"`
}
