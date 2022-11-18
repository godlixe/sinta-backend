package entity

import (
	"sinta-backend/helpers"

	"gorm.io/gorm"
)

type Toko struct {
	ID       uint64 `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"-"`
	BaseModel
}

func (t *Toko) BeforeCreate(tx *gorm.DB) error {
	var err error
	t.Password, err = helpers.HashAndSalt(t.Password)
	if err != nil {
		return err
	}
	return nil
}
