package entity

import (
	"time"

	"gorm.io/gorm"
)

// Base model that includes uint64 ID and created, updated, deleted timestamps
type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type AuthReturn struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
