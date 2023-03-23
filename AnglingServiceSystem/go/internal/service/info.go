package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
	"os"
	"strconv"
	"time"
)

//AddInfo 发布资讯
func AddInfo(c *gin.Context) {
	var info models.Info
	c.ShouldBind(&info)
	info.CreateTime = time.Now()

	token := c.MustGet("token")

	fmt.Println("发布资讯:", info)

	newInfo, err := dao.AddInfo(info)
	fmt.Println(newInfo)

	imagePath := "./web/static/images/newImage.jpg"
	newImage := "./web/static/images/info" + strconv.FormatInt(newInfo.ID, 10) + ".jpg"
	err = os.Rename(imagePath, newImage)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("发布资讯失败！")
	} else {
		fmt.Println("发布资讯成功！")
		c.JSON(200, gin.H{
			"result": "发布资讯成功！",
			"token":  token,
		})
	}
}

// DeleteInfo 删除某条资讯
func DeleteInfo(c *gin.Context) {
	var info models.Info
	c.ShouldBind(&info)
	token := c.MustGet("token")
	fmt.Println("删除资讯：", info)

	err := dao.DeleteInfo(info)

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

// UpdateInfo 修改资讯
func UpdateInfo(c *gin.Context) {
	var info models.Info
	c.ShouldBind(&info)
	token := c.MustGet("token")
	fmt.Println("修改资讯:", info)

	err := dao.UpdateInfo(info)

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

// GetOneInfo 获取一条资讯信息
func GetOneInfo(c *gin.Context) {
	var info models.Info
	c.ShouldBind(&info)
	token := c.MustGet("token")
	fmt.Println("单条资讯：", info.ID)

	info1, err := dao.GetOneInfo(int(info.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(info1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result": "获取成功！",
			"info":   info1,
			"token":  token,
		})
	}
}

// GetAllInfo 获取所有资讯
func GetAllInfo(c *gin.Context) {
	infoList, err := dao.GetAllInfo()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"info_list": infoList,
			"token":     token,
		})
	}
}
