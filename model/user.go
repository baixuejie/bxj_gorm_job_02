package model

type User struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"username" gorm:"unindex" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" validate:"email" binding:"required"`
}
