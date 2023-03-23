package models

import "time"

//Recommend 推荐
type Recommend struct {
	ID            int64     `json:"id" gorm:"column:id"`
	FishName      string    `json:"fish_name" gorm:"column:fish_name"`
	FishInfo      string    `json:"fish_info" gorm:"column:fish_info"`
	RecommendInfo string    `json:"recommend_info" gorm:"column:recommend_info"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}
