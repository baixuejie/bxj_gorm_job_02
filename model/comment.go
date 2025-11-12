package model

import "time"

type Comment struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	PostId    uint      `json:"postId" binding:"required"`
	UserId    uint      `json:"userId" binding:"required"`
}
