package model

import "time"

type Post struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	UserId    uint      `json:"userId" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
