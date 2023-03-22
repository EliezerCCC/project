package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
	"time"
)

//AddNotice 发布公告
func AddNotice(c *gin.Context) {
	var notice models.Notice
	c.ShouldBind(&notice)
	notice.CreateTime = time.Now()

	token := c.MustGet("token")

	fmt.Println("发布公告:", notice)

	err := dao.AddNotice(notice)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("发布公告失败！")
	} else {
		fmt.Println("发布公告成功！")
		c.JSON(200, gin.H{
			"result": "发布公告成功！",
			"token":  token,
		})
	}
}

// DeleteNotice 删除某条公告
func DeleteNotice(c *gin.Context) {
	var notice models.Notice
	c.ShouldBind(&notice)
	token := c.MustGet("token")
	fmt.Println("删除公告：", notice)

	err := dao.DeleteNotice(notice)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("删除失败！")
	} else {
		fmt.Println("删除成功！")
		c.JSON(200, gin.H{
			"result": "删除成功！",
			"token":  token,
		})
	}
}

// UpdateNotice 修改公告
func UpdateNotice(c *gin.Context) {
	var notice models.Notice
	c.ShouldBind(&notice)
	token := c.MustGet("token")
	fmt.Println("修改公告:", notice)

	err := dao.UpdateNotice(notice)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("修改失败！")
	} else {
		fmt.Println("修改成功！")
		c.JSON(200, gin.H{
			"result": "修改成功！",
			"token":  token,
		})
	}

}

// GetOneNotice 获取一条公告信息
func GetOneNotice(c *gin.Context) {
	var notice models.Notice
	c.ShouldBind(&notice)
	token := c.MustGet("token")
	fmt.Println("单条公告：", notice.ID)

	notice1, err := dao.GetOneNotice(int(notice.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(notice1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result": "获取成功！",
			"notice": notice1,
			"token":  token,
		})
	}
}

// GetAllNotice 获取所有公告
func GetAllNotice(c *gin.Context) {
	noticeList, err := dao.GetAllNotice()
	fmt.Println("所有公告信息 ", noticeList)
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"notice_list": noticeList,
			"token":       token,
		})
	}
}
