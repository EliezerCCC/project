package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// AddComment 发布评论
func AddComment(comment models.Comment) (models.Comment, error) {
	err := util.MysqlDB.Create(&comment).Error
	return comment, err
}

// GetComment 获取一个帖子的评论
func GetComment(postID int) (commentList []*models.Comment, err error) {
	err = util.MysqlDB.Where("post_id = ?", postID).Find(&commentList).Error
	return
}
