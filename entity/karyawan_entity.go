package entity

type Karyawan struct {
	ID       uint64 `json:"id"`
	Nama     string `json:"nama"`
	NoTelp   string `json:"no_telp"`
	Alamat   string `json:"alamat"`
	FotoPath string `json:"foto_path"`
	Role     string `json:"role"`
	TokoID   uint64 `json:"toko_id" gorm:"foreignKey"`
	Toko     *Toko  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	BaseModel
}
