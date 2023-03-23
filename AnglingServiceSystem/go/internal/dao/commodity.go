package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddCommodity 发布商品
func AddCommodity(commodity models.Commodity) (models.Commodity, error) {
	err := util.MysqlDB.Create(&commodity).Error
	return commodity, err
}

// GetAllCommodity 所有商品信息
func GetAllCommodity() (commodityList []models.Commodity, err error) {
	err = util.MysqlDB.Find(&commodityList).Error
	return
}

// DeleteCommodity 删除商品
func DeleteCommodity(commodity models.Commodity) (err error) {
	err = util.MysqlDB.Where("Id= ?", commodity.ID).Delete(&models.Commodity{}).Error
	return
}

// UpdateCommodity 修改商品信息
func UpdateCommodity(commodity models.Commodity) (err error) {
	err = util.MysqlDB.Where("Id = ?", commodity.ID).Omit("create_time", "id").Save(&commodity).Error
	return
}

// GetOneCommodity 获取一条商品信息
func GetOneCommodity(commodityID int) (commodity *models.Commodity, err error) {
	err = util.MysqlDB.Where("Id = ?", commodityID).First(&commodity).Error
	return
}
