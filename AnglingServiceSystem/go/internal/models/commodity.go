package models

import "time"

//Commodity 商品
type Commodity struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Introduction string    `json:"introduction" gorm:"column:introduction"` //简介
	Price        float64   `json:"price" gorm:"column:price"`
	Image        string    `json:"image" gorm:"column:image"` //图片路径
	Status       string    `json:"status" gorm:"column:status"`
	SellerID     string    `json:"seller_id" gorm:"column:seller_id"`
	SellerName   string    `json:"seller_name" gorm:"column:seller_name"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
}
