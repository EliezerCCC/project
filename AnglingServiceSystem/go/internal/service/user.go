package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/middleware"
	"go/internal/models"
)

// Login 用户登录
func Login(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)

	fmt.Println("用户登录:", user)

	result, err := dao.Login(user)

	//生成token
	token, err := middleware.GenToken(user)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("登录出错！")
	} else if result == "登录成功" {
		fmt.Println("登录成功！")
		c.JSON(200, gin.H{
			"result": result,
			"token":  token,
		})
	} else {
		fmt.Println("登录失败！")
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}

// Register 用户注册

func Register(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)

	fmt.Println("注册用户:", user)

	result, err := dao.Register(user)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("注册出错！")
	} else {
		fmt.Println("操作注册成功！")
		c.JSON(200, gin.H{"result": result})
	}
}
