package models

import "time"

//Post 帖子
type Post struct {
	ID         int64     `json:"id" gorm:"column:id"`
	Title      string    `json:"title" gorm:"column:title"`
	Content    string    `json:"content" gorm:"column:content"`
	Author     string    `json:"author" gorm:"column:author"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
