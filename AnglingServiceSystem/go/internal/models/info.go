package models

import "time"

//Info 资讯
type Info struct {
	Title      string    `json:"title" gorm:"column:title"`
	Content    string    `json:"content" gorm:"column:content"`
	createTime time.Time `json:"create_time" gorm:"column:create_time"`
}
