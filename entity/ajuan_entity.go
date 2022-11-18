package entity

type Ajuan struct {
	ID          uint64        `json:"id"`
	Status      bool          `json:"status"`
	TokoID      uint64        `json:"toko_id" gorm:"foreignKey"`
	Toko        *Toko         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
	DetailAjuan []DetailAjuan `json:"detail_ajuan"`
	BaseModel
}
