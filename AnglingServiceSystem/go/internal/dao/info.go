package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddInfo 发布资讯
func AddInfo(info models.Info) (err error) {
	err = util.MysqlDB.Create(&info).Error
	return
}

// GetAllInfo 所有资讯信息
func GetAllInfo() (infoList []models.Info, err error) {
	err = util.MysqlDB.Find(&infoList).Error
	return
}

// DeleteInfo 删除资讯
func DeleteInfo(info models.Info) (err error) {
	err = util.MysqlDB.Where("Id= ?", info.ID).Delete(&models.Info{}).Error
	return
}

// UpdateInfo 修改资讯信息
func UpdateInfo(info models.Info) (err error) {
	err = util.MysqlDB.Where("Id = ?", info.ID).Save(&info).Error
	return
}

// GetOneInfo 获取一条资讯信息
func GetOneInfo(infoID int) (info *models.Info, err error) {
	err = util.MysqlDB.Where("Id = ?", infoID).First(&info).Error
	return
}
