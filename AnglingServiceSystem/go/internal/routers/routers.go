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

	//用户登录
	r.POST("/login", service.Login)
	//用户注册
	r.POST("/register", service.Register)

	//获取所有用户信息
	r.GET("/getAllUser", middleware.JWTAuthMiddleware(), service.GetAllUser)
	//获取某个用户信息
	r.POST("/getOneUser", middleware.JWTAuthMiddleware(), service.GetOneUser)
	//修改某个用户信息
	r.POST("/updateUser", middleware.JWTAuthMiddleware(), service.UpdateUser)

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

	//发布商品
	r.POST("/addCommodity", middleware.JWTAuthMiddleware(), service.AddCommodity)
	//获取所有商品
	r.GET("/getAllCommodity", middleware.JWTAuthMiddleware(), service.GetAllCommodity)
	//获取某件商品
	r.POST("/getOneCommodity", middleware.JWTAuthMiddleware(), service.GetOneCommodity)
	//修改某件商品
	r.POST("/updateCommodity", middleware.JWTAuthMiddleware(), service.UpdateCommodity)
	//删除商品
	r.POST("/deleteCommodity", middleware.JWTAuthMiddleware(), service.DeleteCommodity)

	//发布钓场
	r.POST("/addAnglingSite", middleware.JWTAuthMiddleware(), service.AddAnglingSite)
	//获取所有钓场
	r.GET("/getAllAnglingSite", middleware.JWTAuthMiddleware(), service.GetAllAnglingSite)
	//获取某钓场
	r.POST("/getOneAnglingSite", middleware.JWTAuthMiddleware(), service.GetOneAnglingSite)
	//修改某钓场
	r.POST("/updateAnglingSite", middleware.JWTAuthMiddleware(), service.UpdateAnglingSite)
	//删除钓场
	r.POST("/deleteAnglingSite", middleware.JWTAuthMiddleware(), service.DeleteAnglingSite)

	//添加推荐
	r.POST("/addRecommend", middleware.JWTAuthMiddleware(), service.AddRecommend)
	//获取所有推荐
	r.GET("/getAllRecommend", middleware.JWTAuthMiddleware(), service.GetAllRecommend)
	//获取某推荐
	r.POST("/getOneRecommend", middleware.JWTAuthMiddleware(), service.GetOneRecommend)
	//修改某推荐
	r.POST("/updateRecommend", middleware.JWTAuthMiddleware(), service.UpdateRecommend)
	//删除推荐
	r.POST("/deleteRecommend", middleware.JWTAuthMiddleware(), service.DeleteRecommend)

	//发布帖子
	r.POST("/addPost", middleware.JWTAuthMiddleware(), service.AddPost)
	//获取所有帖子
	r.GET("/getAllPost", middleware.JWTAuthMiddleware(), service.GetAllPost)
	//获取某帖子
	r.POST("/getOnePost", middleware.JWTAuthMiddleware(), service.GetOnePost)
	//修改某帖子
	r.POST("/updatePost", middleware.JWTAuthMiddleware(), service.UpdatePost)
	//删除帖子
	r.POST("/deletePost", middleware.JWTAuthMiddleware(), service.DeletePost)

	//上传图片
	r.POST("/upload", service.Upload)
	//获取图片
	r.GET("/getImage", service.Image)

	return r
}
