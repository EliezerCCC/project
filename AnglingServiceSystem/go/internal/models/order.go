package models

import "time"

//Order 订单
type Order struct {
	ID          int64     `json:"id" gorm:"column:id"`
	UserID      string    `json:"uer_id" gorm:"column:user_id"`
	CommodityID int64     `json:"commodity_id" gorm:"column:commodity_id"`
	Amount      float64   `json:"amount" gorm:"column:amount"`
	Status      string    `json:"status" gorm:"column:status"`
	Address     Address   `json:"address" gorm:"column:address"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}
