package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddAnglingSite 发布钓场
func AddAnglingSite(anglingSite models.AnglingSite) (models.AnglingSite, error) {
	err := util.MysqlDB.Create(&anglingSite).Error
	return anglingSite, err
}

// GetAllAnglingSite 所有钓场信息
func GetAllAnglingSite() (anglingSiteList []models.AnglingSite, err error) {
	err = util.MysqlDB.Find(&anglingSiteList).Error
	return
}

// DeleteAnglingSite 删除钓场
func DeleteAnglingSite(anglingSite models.AnglingSite) (err error) {
	err = util.MysqlDB.Where("Id= ?", anglingSite.ID).Delete(&models.AnglingSite{}).Error
	return
}

// UpdateAnglingSite 修改钓场信息
func UpdateAnglingSite(anglingSite models.AnglingSite) (err error) {
	err = util.MysqlDB.Where("Id = ?", anglingSite.ID).Omit("create_time", "id").Save(&anglingSite).Error
	return
}

// GetOneAnglingSite 获取一个钓场信息
func GetOneAnglingSite(anglingSiteID int) (anglingSite *models.AnglingSite, err error) {
	err = util.MysqlDB.Where("Id = ?", anglingSiteID).First(&anglingSite).Error
	return
}
