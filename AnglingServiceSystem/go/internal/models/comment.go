package models

import "time"

//Comment 评论
type Comment struct {
	ID         int64     `json:"id"`
	PostID     int64     `json:"post_id"`
	createTime time.Time `json:"create_time" gorm:"column:create_time"`
}
