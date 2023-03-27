package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
)

//AddAddress 添加地址
func AddAddress(c *gin.Context) {
	var address models.Address
	c.ShouldBind(&address)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	address.UserID = userIDValue

	token := c.MustGet("token")

	fmt.Println("添加地址:", address)

	newAddress, err := dao.AddAddress(address)
	fmt.Println(newAddress)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("添加地址失败！")
	} else {
		fmt.Println("添加地址成功！")
		c.JSON(200, gin.H{
			"result": "添加地址成功！",
			"token":  token,
		})
	}
}

//GetAddress 获取地址
func GetAddress(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)

	addressList, err := dao.GetAddress(user.ID)
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"address_list": addressList,
			"token":        token,
		})
	}
}

// DeleteAddress 删除地址
func DeleteAddress(c *gin.Context) {
	var address models.Address
	c.ShouldBind(&address)
	token := c.MustGet("token")
	fmt.Println("删除地址：", address)

	err := dao.DeleteAddress(address)

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
