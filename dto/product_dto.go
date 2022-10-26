package dto

type ProductCreateDTO struct {
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required" binding:"required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required" binding:"required"`
	UserID      uint   `json:"userId" form:"userId" binding:"required"`
}

type ProductUpdateDTO struct {
	ID          uint
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required" binding:"required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required" binding:"required"`
	UserID      uint   `json:"userId" form:"userId" binding:"required"`
}
