package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddPost 发布帖子
func AddPost(post models.Post) (models.Post, error) {
	err := util.MysqlDB.Create(&post).Error
	return post, err
}

// GetAllPost 所有帖子信息
func GetAllPost() (postList []models.Post, err error) {
	err = util.MysqlDB.Order("create_time desc").Find(&postList).Error
	return
}

// DeletePost 删除帖子
func DeletePost(post models.Post) (err error) {
	err = util.MysqlDB.Where("Id= ?", post.ID).Delete(&models.Post{}).Error
	return
}

// UpdatePost 修改帖子信息
func UpdatePost(post models.Post) (err error) {
	err = util.MysqlDB.Where("Id = ?", post.ID).Omit("create_time", "id").Save(&post).Error
	return
}

// GetOnePost 获取一条帖子信息
func GetOnePost(postID int) (post *models.Post, err error) {
	err = util.MysqlDB.Where("Id = ?", postID).First(&post).Error
	return
}
