package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddRecommend 发布推荐
func AddRecommend(recommend models.Recommend) (models.Recommend, error) {
	err := util.MysqlDB.Create(&recommend).Error
	return recommend, err
}

// GetAllRecommend 所有推荐信息
func GetAllRecommend() (recommendList []models.Recommend, err error) {
	err = util.MysqlDB.Find(&recommendList).Error
	return
}

// DeleteRecommend 删除推荐
func DeleteRecommend(recommend models.Recommend) (err error) {
	err = util.MysqlDB.Where("Id= ?", recommend.ID).Delete(&models.Recommend{}).Error
	return
}

// UpdateRecommend 修改推荐信息
func UpdateRecommend(recommend models.Recommend) (err error) {
	err = util.MysqlDB.Where("Id = ?", recommend.ID).Omit("create_time", "id").Save(&recommend).Error
	return
}

// GetOneRecommend 获取一条推荐信息
func GetOneRecommend(recommendID int) (recommend *models.Recommend, err error) {
	err = util.MysqlDB.Where("Id = ?", recommendID).First(&recommend).Error
	return
}
