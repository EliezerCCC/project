package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddAddress 添加地址
func AddAddress(address models.Address) (models.Address, error) {
	err := util.MysqlDB.Create(&address).Error
	return address, err
}

// GetAddress 获取一个用户的地址
func GetAddress(userID string) (addressList []*models.Address, err error) {
	err = util.MysqlDB.Where("user_id = ?", userID).Find(&addressList).Error
	return
}

// DeleteAddress 删除地址
func DeleteAddress(address models.Address) (err error) {
	err = util.MysqlDB.Where("id = ?", address.ID).Delete(&models.Address{}).Error
	return
}
