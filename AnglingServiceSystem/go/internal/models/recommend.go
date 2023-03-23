package models

//Recommend 推荐
type Recommend struct {
	ID            int64  `json:"id" gorm:"column:id"`
	FishName      string `json:"fish_name" gorm:"column:fish_name"`
	FishInfo      string `json:"fish_info" gorm:"column:fish_info"`
	RecommendInfo string `json:"recommend_info" gorm:"column:recommend_info"`
}
