package models

//User 用户
type User struct {
	ID           string `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	Password     string `json:"password" gorm:"column:password"`
	Identity     string `json:"identity" gorm:"column:identity"`         //身份
	Introduction string `json:"introduction" gorm:"column:introduction"` //简介
	Status       string `json:"status" gorm:"column:status"`             //状态
}
