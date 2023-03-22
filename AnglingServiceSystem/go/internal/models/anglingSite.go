package models

import "time"

//AnglingSite 钓场
type AnglingSite struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Place      string    `json:"place"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
