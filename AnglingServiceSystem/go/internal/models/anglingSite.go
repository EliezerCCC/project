package models

import "time"

//AnglingSite 钓场
type AnglingSite struct {
	ID           int64     `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Address      string    `json:"address" gorm:"column:address"`
	Introduction string    `json:"introduction" gorm:"column:introduction"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
}
