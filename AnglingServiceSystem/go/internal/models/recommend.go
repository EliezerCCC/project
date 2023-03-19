package models

//Recommend 推荐
type Recommend struct {
	FishName      string `json:"fish_name" gorm:"column:fish_name"`
	FishInfo      string `json:"fish_info" gorm:"column:fish_info"`
	RecommendInfo string `json:"recommend_info" gorm:"column:recommend_info"`
}
