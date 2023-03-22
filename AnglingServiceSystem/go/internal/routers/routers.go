package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/middleware"
	"go/internal/service"
	"net/http"
	"path"
)

func SetupRouter() *gin.Engine {
	//创建默认路由引擎
	r := gin.Default()

	//允许跨域访问
	r.Use(middleware.Cors())

	//用户登录
	r.POST("/login", service.Login)
	//用户注册
	r.POST("/register", service.Register)

	//发布公告
	r.POST("/addNotice", middleware.JWTAuthMiddleware(), service.AddNotice)
	//获取所有公告
	r.GET("/getAllNotice", middleware.JWTAuthMiddleware(), service.GetAllNotice)
	//获取某条公告
	r.POST("/getOneNotice", middleware.JWTAuthMiddleware(), service.GetOneNotice)
	//修改某条公告
	r.POST("/updateNotice", middleware.JWTAuthMiddleware(), service.UpdateNotice)
	//删除公告
	r.POST("/deleteNotice", middleware.JWTAuthMiddleware(), service.DeleteNotice)

	//发布资讯
	r.POST("/addInfo", middleware.JWTAuthMiddleware(), service.AddInfo)
	//获取所有资讯
	r.GET("/getAllInfo", middleware.JWTAuthMiddleware(), service.GetAllInfo)
	//获取某条资讯
	r.POST("/getOneInfo", middleware.JWTAuthMiddleware(), service.GetOneInfo)
	//修改某条资讯
	r.POST("/updateInfo", middleware.JWTAuthMiddleware(), service.UpdateInfo)
	//删除资讯
	r.POST("/deleteInfo", middleware.JWTAuthMiddleware(), service.DeleteInfo)

	//获取所有商品
	r.GET("/getAllCommodity", service.GetAllCommodity)

	r.POST("/api/private/v1/upload", func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "接收文件失败")
			return
		}
		fmt.Println(f.Filename)
		dst := path.Join("./web/static/images", "tupian2.jpg")
		if err := c.SaveUploadedFile(f, dst); err != nil {
			c.String(http.StatusBadRequest, "保存文件失败")
			return
		}
		c.String(http.StatusOK, "上传文件成功")

	})

	return r
}
