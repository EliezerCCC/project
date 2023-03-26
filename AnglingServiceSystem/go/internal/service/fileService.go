package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path"
)

func Upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "接收文件失败")
		return
	}

	dst := path.Join("./web/static/images", "newImage.jpg")
	if err := c.SaveUploadedFile(f, dst); err != nil {
		c.String(http.StatusBadRequest, "保存文件失败")
		return
	}
	c.String(http.StatusOK, "上传文件成功")

}

func Image(c *gin.Context) {
	//显示图片的方法
	imageName := c.Query("imageName")       //截取get请求参数，也就是图片的路径，可是使用绝对路径，也可使用相对路径
	file, err := ioutil.ReadFile(imageName) //把要显示的图片读取到变量中
	if err != nil {
		file, _ = ioutil.ReadFile("./web/static/images/tupian.jpg")
	}
	c.Writer.WriteString(string(file)) //关键一步，写给前端
}
