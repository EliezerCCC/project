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

//AddRecommend 添加推荐
func AddRecommend(c *gin.Context) {
	var recommend models.Recommend
	c.ShouldBind(&recommend)
	recommend.CreateTime = time.Now()
	token := c.MustGet("token")
	fmt.Println("添加推荐:", recommend)

	newRecommend, err := dao.AddRecommend(recommend)
	fmt.Println(newRecommend)

	imagePath := "./web/static/images/newImage.jpg"
	newImage := "./web/static/images/recommend" + strconv.FormatInt(newRecommend.ID, 10) + ".jpg"
	err = os.Rename(imagePath, newImage)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("添加推荐失败！")
	} else {
		fmt.Println("添加推荐成功！")
		c.JSON(200, gin.H{
			"result": "添加推荐成功！",
			"token":  token,
		})
	}
}

// DeleteRecommend 删除某推荐
func DeleteRecommend(c *gin.Context) {
	var recommend models.Recommend
	c.ShouldBind(&recommend)
	token := c.MustGet("token")
	fmt.Println("删除推荐：", recommend)

	err := dao.DeleteRecommend(recommend)

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

// UpdateRecommend 修改推荐
func UpdateRecommend(c *gin.Context) {
	var recommend models.Recommend
	c.ShouldBind(&recommend)
	token := c.MustGet("token")
	fmt.Println("修改推荐:", recommend)

	err := dao.UpdateRecommend(recommend)

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

// GetOneRecommend 获取一条推荐信息
func GetOneRecommend(c *gin.Context) {
	var recommend models.Recommend
	c.ShouldBind(&recommend)
	token := c.MustGet("token")
	fmt.Println("单条推荐：", recommend.ID)

	recommend1, err := dao.GetOneRecommend(int(recommend.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(recommend1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result":    "获取成功！",
			"recommend": recommend1,
			"token":     token,
		})
	}
}

// GetAllRecommend 获取所有推荐
func GetAllRecommend(c *gin.Context) {
	recommendList, err := dao.GetAllRecommend()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"recommend_list": recommendList,
			"token":          token,
		})
	}
}
