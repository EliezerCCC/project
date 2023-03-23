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

//AddAnglingSite 发布钓场
func AddAnglingSite(c *gin.Context) {
	var anglingSite models.AnglingSite
	c.ShouldBind(&anglingSite)
	anglingSite.CreateTime = time.Now()
	token := c.MustGet("token")
	fmt.Println("发布钓场:", anglingSite)

	newAnglingSite, err := dao.AddAnglingSite(anglingSite)
	fmt.Println(newAnglingSite)

	imagePath := "./web/static/images/newImage.jpg"
	newImage := "./web/static/images/anglingSite" + strconv.FormatInt(newAnglingSite.ID, 10) + ".jpg"
	err = os.Rename(imagePath, newImage)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("发布钓场失败！")
	} else {
		fmt.Println("发布钓场成功！")
		c.JSON(200, gin.H{
			"result": "发布钓场成功！",
			"token":  token,
		})
	}
}

// DeleteAnglingSite 删除某钓场
func DeleteAnglingSite(c *gin.Context) {
	var anglingSite models.AnglingSite
	c.ShouldBind(&anglingSite)
	token := c.MustGet("token")
	fmt.Println("删除钓场：", anglingSite)

	err := dao.DeleteAnglingSite(anglingSite)

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

// UpdateAnglingSite 修改钓场
func UpdateAnglingSite(c *gin.Context) {
	var anglingSite models.AnglingSite
	c.ShouldBind(&anglingSite)
	token := c.MustGet("token")
	fmt.Println("修改钓场:", anglingSite)

	err := dao.UpdateAnglingSite(anglingSite)

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

// GetOneAnglingSite 获取一个钓场信息
func GetOneAnglingSite(c *gin.Context) {
	var anglingSite models.AnglingSite
	c.ShouldBind(&anglingSite)
	token := c.MustGet("token")
	fmt.Println("单个钓场：", anglingSite.ID)

	anglingSite1, err := dao.GetOneAnglingSite(int(anglingSite.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(anglingSite1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result":      "获取成功！",
			"anglingSite": anglingSite1,
			"token":       token,
		})
	}
}

// GetAllAnglingSite 获取所有钓场
func GetAllAnglingSite(c *gin.Context) {
	anglingSiteList, err := dao.GetAllAnglingSite()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"anglingSite_list": anglingSiteList,
			"token":            token,
		})
	}
}
