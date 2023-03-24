package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddNotice 发布公告
func AddNotice(notice models.Notice) (err error) {
	err = util.MysqlDB.Create(&notice).Error
	return
}

// GetAllNotice 所有公告信息
func GetAllNotice() (noticeList []models.Notice, err error) {
	err = util.MysqlDB.Order("create_time desc").Find(&noticeList).Error
	return
}

// DeleteNotice 删除公告
func DeleteNotice(notice models.Notice) (err error) {
	err = util.MysqlDB.Where("Id= ?", notice.ID).Delete(&models.Notice{}).Error
	return
}

// UpdateNotice 修改公告信息
func UpdateNotice(notice models.Notice) (err error) {
	err = util.MysqlDB.Where("Id = ?", notice.ID).Omit("create_time", "id").Save(&notice).Error
	return
}

// GetOneNotice 获取一条公告信息
func GetOneNotice(noticeID int) (notice *models.Notice, err error) {
	err = util.MysqlDB.Where("Id = ?", noticeID).First(&notice).Error
	return
}
