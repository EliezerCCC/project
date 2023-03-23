package models

import "time"

//Commodity 商品
type Commodity struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Introduction string    `json:"introduction" gorm:"column:introduction"`
	Amount       int       `json:"amount" gorm:"column:amount"`
	Type         string    `json:"type" gorm:"column:type"`
	Price        float64   `json:"price" gorm:"column:price"`
	Status       string    `json:"status" gorm:"column:status"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
}
