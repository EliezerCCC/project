package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/internal/dao"
	"go/internal/models"
	"os"
	"strconv"
	"time"
)

//AddPost 发布帖子
func AddPost(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)
	post.CreateTime = time.Now()
	token := c.MustGet("token")

	userName, _ := c.Get("user_name")
	userID, _ := c.Get("user_id")

	userNameValue, _ := userName.(string)
	userIDValue, _ := userID.(string)

	post.UserName = userNameValue
	post.UserID = userIDValue

	fmt.Println("发布帖子:", post)

	newPost, err := dao.AddPost(post)
	fmt.Println(newPost)

	imagePath := "./web/static/images/newImage.jpg"
	newImage := "./web/static/images/post" + strconv.FormatInt(newPost.ID, 10) + ".jpg"
	err = os.Rename(imagePath, newImage)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error(), "token": token})
		fmt.Println("发布帖子失败！")
	} else {
		fmt.Println("发布帖子成功！")
		c.JSON(200, gin.H{
			"result": "发布帖子成功！",
			"token":  token,
		})
	}
}

// DeletePost 删除某帖子
func DeletePost(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)
	token := c.MustGet("token")
	fmt.Println("删除帖子：", post)

	err := dao.DeletePost(post)

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

// UpdatePost 修改帖子
func UpdatePost(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)
	token := c.MustGet("token")
	fmt.Println("修改帖子:", post)

	err := dao.UpdatePost(post)

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

// GetOnePost 获取一条帖子信息
func GetOnePost(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)
	token := c.MustGet("token")
	fmt.Println("单条帖子：", post.ID)

	post1, err := dao.GetOnePost(int(post.ID))

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println(post1)
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result": "获取成功！",
			"post":   post1,
			"token":  token,
		})
	}
}

// GetAllPost 获取所有帖子
func GetAllPost(c *gin.Context) {
	postList, err := dao.GetAllPost()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"post_list": postList,
			"token":     token,
		})
	}
}
