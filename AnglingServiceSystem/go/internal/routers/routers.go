package routers

import (
	"github.com/gin-gonic/gin"
	"go/internal/middleware"
	"go/internal/service"
)

func SetupRouter() *gin.Engine {
	//创建默认路由引擎
	r := gin.Default()

	//允许跨域访问
	r.Use(middleware.Cors())

	//获取所有商品
	r.GET("/getAllCommodity", service.GetAllCommodity)

	return r
}
