package entity

type Comment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	UserID  uint64 `json:"user_id"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	PhotoID uint64 `json:"photo_id"`
	Photo   Photo  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Message string `json:"message"`
	BaseModel
}
