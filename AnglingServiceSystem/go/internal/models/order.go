package models

import "time"

//Order 订单
type Order struct {
	ID          int64     `json:"id"`
	BuyerID     string    `json:"buyer_id"`
	SellerID    string    `json:"seller_id"`
	CommodityID int64     `json:"commodity_id"`
	Price       float64   `json:"price"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}
