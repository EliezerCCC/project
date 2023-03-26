package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
	"time"
)

//AddComment 发布评论
func AddComment(c *gin.Context) {
	var comment models.Comment
	c.ShouldBind(&comment)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	userName, _ := c.Get("user_name")
	userNameValue, _ := userName.(string)

	comment.CreateTime = time.Now()
	comment.UserID = userIDValue
	comment.UserName = userNameValue

	token := c.MustGet("token")

	fmt.Println("发布评论:", comment)

	newInfo, err := dao.AddComment(comment)
	fmt.Println(newInfo)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("发布评论失败！")
	} else {
		fmt.Println("发布评论成功！")
		c.JSON(200, gin.H{
			"result": "发布评论成功！",
			"token":  token,
		})
	}
}

//GetComment 获取评论
func GetComment(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)

	commentList, err := dao.GetComment(int(post.ID))
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"comment_list": commentList,
			"token":        token,
		})
	}
}
