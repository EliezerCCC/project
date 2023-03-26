package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddCollect 添加收藏
func AddCollect(collect models.Collect) (models.Collect, error) {
	err := util.MysqlDB.Create(&collect).Error
	return collect, err
}

// DeleteCollect 删除收藏
func DeleteCollect(collect models.Collect) (err error) {
	err = util.MysqlDB.Where("user_id = ? AND info_id = ?", collect.UserID, collect.InfoID).Delete(&models.Collect{}).Error
	return
}

// IsCollect 查询是否收藏
func IsCollect(collect models.Collect) (bool, error) {
	result := util.MysqlDB.Where("user_id = ? AND info_id = ?", collect.UserID, collect.InfoID).First(&models.Collect{})
	row := result.RowsAffected
	if row != 0 {
		return true, result.Error
	}
	return false, result.Error
}

// PersonCollect 查询个人收藏
func PersonCollect(userID string) (collectList []models.Collect, err error) {
	err = util.MysqlDB.Where("user_id = ?", userID).Find(&collectList).Error
	return
}
