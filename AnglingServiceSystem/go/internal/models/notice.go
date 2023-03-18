package models

import "time"

//Notice 公告
type Notice struct {
	Title      string    `json:"title" gorm:"column:title"`
	Content    string    `json:"content" gorm:"column:content"`
	createTime time.Time `json:"create_time" gorm:"column:create_time"`
}
