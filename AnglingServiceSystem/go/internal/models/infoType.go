package models

//InfoType 资讯类型
type InfoType struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}
