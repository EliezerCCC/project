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

	identity, name, result, err := dao.Login(user)
	user.Name = name

	//生成token
	token, err := middleware.GenToken(user)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("登录出错！")
	} else if result == "登录成功" {
		fmt.Println("登录成功！")
		c.JSON(200, gin.H{
			"result":   result,
			"identity": identity,
			"token":    token,
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

// UpdateUser 修改用户信息
func UpdateUser(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	token := c.MustGet("token")
	fmt.Println("修改用户:", user)

	err := dao.UpdateUser(user)

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

// GetOneUser 获取一个用户信息
func GetOneUser(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	token := c.MustGet("token")
	fmt.Println("单条推荐：", user.ID)

	user1, err := dao.GetOneUser(user.ID)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(user1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result": "获取成功！",
			"user":   user1,
			"token":  token,
		})
	}
}

// GetAllUser 获取所有用户信息
func GetAllUser(c *gin.Context) {
	userList, err := dao.GetAllUser()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"user_list": userList,
			"token":     token,
		})
	}
}
