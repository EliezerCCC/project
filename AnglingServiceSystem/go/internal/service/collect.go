package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
)

//AddCollect 添加收藏
func AddCollect(c *gin.Context) {
	var collect models.Collect
	c.ShouldBind(&collect)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	collect.UserID = userIDValue

	token := c.MustGet("token")
	fmt.Println("添加收藏:", collect)

	_, err := dao.AddCollect(collect)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("添加收藏失败！")
	} else {
		fmt.Println("添加收藏成功！")
		c.JSON(200, gin.H{
			"result": "添加收藏成功！",
			"token":  token,
		})
	}
}

//DeleteCollect 删除收藏
func DeleteCollect(c *gin.Context) {
	var collect models.Collect
	c.ShouldBind(&collect)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	collect.UserID = userIDValue

	token := c.MustGet("token")
	fmt.Println("删除收藏:", collect)

	err := dao.DeleteCollect(collect)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("删除收藏失败！")
	} else {
		fmt.Println("删除收藏成功！")
		c.JSON(200, gin.H{
			"result": "删除收藏成功！",
			"token":  token,
		})
	}
}

//IsCollect 查询收藏
func IsCollect(c *gin.Context) {
	var collect models.Collect
	c.ShouldBind(&collect)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	collect.UserID = userIDValue

	token := c.MustGet("token")

	is, err := dao.IsCollect(collect)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("查询收藏失败！", err)
	} else {
		fmt.Println("查询收藏成功！")
		c.JSON(200, gin.H{
			"is_collect": is,
			"result":     "查询收藏成功！",
			"token":      token,
		})
	}
}

// PersonCollect 获取个人收藏
func PersonCollect(c *gin.Context) {

	infoList := []*models.Info{}

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)

	collectList, err := dao.PersonCollect(userIDValue)

	for _, v := range collectList {
		info1, err := dao.GetOneInfo(int(v.InfoID))
		if err == nil {
			infoList = append(infoList, info1)
		}
	}

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
