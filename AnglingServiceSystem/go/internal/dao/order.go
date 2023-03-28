package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddOrder 新订单
func AddOrder(order models.Order) (models.Order, error) {
	err := util.MysqlDB.Create(&order).Error
	return order, err
}

// GetAllOrder 所有订单信息
func GetAllOrder() (orderList []models.Order, err error) {
	err = util.MysqlDB.Order("create_time desc").Find(&orderList).Error
	return
}

// DeleteOrder 删除订单
func DeleteOrder(order models.Order) (err error) {
	err = util.MysqlDB.Where("Id= ?", order.ID).Delete(&models.Order{}).Error
	return
}

// UpdateOrder 修改订单信息
func UpdateOrder(order models.Order) (err error) {
	err = util.MysqlDB.Where("Id = ?", order.ID).Omit("create_time", "id").Save(&order).Error
	return
}

// GetOrder 获取某用户订单
func GetOrder(userID string) (orderList []models.Order, err error) {
	err = util.MysqlDB.Where("user_id = ?", userID).Order("create_time desc").Find(&orderList).Error
	return
}
