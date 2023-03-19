package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var MysqlDB *gorm.DB

func InitMysql() (err error) {
	//连接mysql
	dsn := "root:123456@tcp(127.0.0.1:3306)/anglingservicesystem?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		}})

	return err
}
