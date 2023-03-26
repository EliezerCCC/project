package models

//Collect 收藏
type Collect struct {
	ID     int64  `json:"id" gorm:"column:id"`
	InfoID int64  `json:"info_id" gorm:"column:info_id"`
	UserID string `json:"user_id" gorm:"column:user_id"`
}
