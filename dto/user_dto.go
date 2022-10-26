package dto

// type UserRegisterDTO struct {
// 	FullName string `gorm:"not null" json:"full_name" form:"full_name" binding:"required"`
// 	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" binding:"required"`
// 	Password string `gorm:"not null" json:"password" binding:"required"`
// 	// Products []entity.Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL, OnDelete:SET NULL;" json:"products"`
// }

type UserRegisterDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserUpdateDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
