package models

import "time"

//Info 资讯
type Info struct {
	ID         int64     `json:"id" gorm:"column:id"`
	Title      string    `json:"title" gorm:"column:title"`
	Content    string    `json:"content" gorm:"column:content"`
	Type       string    `json:"type" gorm:"column:type"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
