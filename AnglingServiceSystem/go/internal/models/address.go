package models

//Address 地址
type Address struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Name     string `json:"name" gorm:"column:name"`
	Province string `json:"province" gorm:"column:province"`
	City     string `json:"city" gorm:"column:city"`
	Area     string `json:"area" gorm:"column:area"`
	Detail   string `json:"detail" gorm:"column:detail"`
	UserID   string `json:"user_id" gorm:"column:user_id"`
}
