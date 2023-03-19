package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
)

// GetAllCommodity 获取所有商品
func GetAllCommodity(c *gin.Context) {
	CommodityList, err := dao.GetAllCommodity()
	fmt.Println("所有商品信息 ", CommodityList)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"commodity_list": CommodityList})
	}
}
