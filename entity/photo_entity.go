package entity

type Photo struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	BaseModel
}
