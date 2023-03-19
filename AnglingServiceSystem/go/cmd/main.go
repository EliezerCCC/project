package main

import (
	"go/internal/routers"
	"go/internal/util"
)

func main() {
	//连接mysql数据库
	err := util.InitMysql()
	if err != nil {
		panic(err)
	}

	//设置接口路由
	r := routers.SetupRouter()

	//启动http服务
	err = r.Run(":9090")
	if err != nil {
		return
	}
}
