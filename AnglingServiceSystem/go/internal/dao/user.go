package dao

import (
	"go/internal/models"
	"go/internal/util"
)

// 用户登录 Login
func Login(user1 models.User) (identity string, name string, result string, err error) {
	user := user1
	err = util.MysqlDB.First(&user).Error

	if user.Password == user1.Password {
		identity = user.Identity
		name = user.Name
		result = "登录成功"
	} else {
		result = "账号或密码错误"
	}

	return
}

// 用户注册 Register

func Register(user models.User) (result string, err error) {
	res := util.MysqlDB.First(&user)
	if res.RowsAffected == 1 {
		result = "账号已存在"
	} else {
		err = util.MysqlDB.Create(&user).Error
		result = "注册信息可用"
	}

	return
}

// GetAllUser 所有用户信息
func GetAllUser() (userList []models.User, err error) {
	err = util.MysqlDB.Find(&userList).Error
	return
}

// UpdateUser 修改用户信息
func UpdateUser(user models.User) (err error) {
	err = util.MysqlDB.Where("Id = ?", user.ID).Omit("id").Save(&user).Error
	return
}

// GetOneUser 获取一个用户信息
func GetOneUser(userID string) (user *models.User, err error) {
	err = util.MysqlDB.Where("Id = ?", userID).First(&user).Error
	return
}
