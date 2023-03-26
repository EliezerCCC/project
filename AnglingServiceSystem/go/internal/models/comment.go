package models

import "time"

//Comment 评论
type Comment struct {
	ID         int64     `json:"id" gorm:"column:id"`
	Content    string    `json:"content" gorm:"column:content"`
	PostID     int64     `json:"post_id" gorm:"column:post_id"`
	UserID     string    `json:"user_id" gorm:"column:user_id"`
	UserName   string    `json:"user_name" gorm:"column:user_name"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
