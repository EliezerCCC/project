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

//AddCommodity 发布商品
func AddCommodity(c *gin.Context) {
	var commodity models.Commodity
	c.ShouldBind(&commodity)
	commodity.CreateTime = time.Now()
	commodity.Status = "下架"
	token := c.MustGet("token")
	fmt.Println("发布商品:", commodity)

	newCommodity, err := dao.AddCommodity(commodity)
	fmt.Println(newCommodity)

	imagePath := "./web/static/images/newImage.jpg"
	newImage := "./web/static/images/commodity" + strconv.FormatInt(newCommodity.ID, 10) + ".jpg"
	err = os.Rename(imagePath, newImage)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("发布商品失败！")
	} else {
		fmt.Println("发布商品成功！")
		c.JSON(200, gin.H{
			"result": "发布商品成功！",
			"token":  token,
		})
	}
}

// DeleteCommodity 删除某件商品
func DeleteCommodity(c *gin.Context) {
	var commodity models.Commodity
	c.ShouldBind(&commodity)
	token := c.MustGet("token")
	fmt.Println("删除商品：", commodity)

	err := dao.DeleteCommodity(commodity)

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

// UpdateCommodity 修改商品
func UpdateCommodity(c *gin.Context) {
	var commodity models.Commodity
	c.ShouldBind(&commodity)
	token := c.MustGet("token")
	fmt.Println("修改商品:", commodity)

	err := dao.UpdateCommodity(commodity)

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

// GetOneCommodity 获取一件商品信息
func GetOneCommodity(c *gin.Context) {
	var commodity models.Commodity
	c.ShouldBind(&commodity)
	token := c.MustGet("token")
	fmt.Println("单件商品：", commodity.ID)

	commodity1, err := dao.GetOneCommodity(int(commodity.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(commodity1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result":    "获取成功！",
			"commodity": commodity1,
			"token":     token,
		})
	}
}

// GetAllCommodity 获取所有商品
func GetAllCommodity(c *gin.Context) {
	commodityList, err := dao.GetAllCommodity()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"commodity_list": commodityList,
			"token":          token,
		})
	}
}
