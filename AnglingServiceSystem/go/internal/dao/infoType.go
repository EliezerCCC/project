package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddInfoType 添加资讯分类
func AddInfoType(infoType models.InfoType) (err error) {
	err = util.MysqlDB.Create(&infoType).Error
	return
}

// GetAllInfoType 所有资讯分类信息
func GetAllInfoType() (infoTypeList []models.InfoType, err error) {
	err = util.MysqlDB.Find(&infoTypeList).Error
	return
}

// DeleteInfoType 删除资讯分类
func DeleteInfoType(infoType models.InfoType) (err error) {
	err = util.MysqlDB.Where("Id= ?", infoType.ID).Delete(&models.InfoType{}).Error
	return
}

// UpdateInfoType 修改资讯分类信息
func UpdateInfoType(infoType models.InfoType) (err error) {
	err = util.MysqlDB.Where("Id = ?", infoType.ID).Save(&infoType).Error
	return
}
