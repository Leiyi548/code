package user

import (
	"fmt"
	"ginStudy/src/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// extened
	BaseController
}

// 大写是因为我们在外部包需要调用它
func (userController UserController) Index(c *gin.Context) {
	c.String(http.StatusOK, "用户首页")
	userController.success(c)
}

func (userController UserController) List(c *gin.Context) {
	c.String(http.StatusOK, "用户列表")
	userController.success(c)
}

func (userController UserController) Description(c *gin.Context) {
	c.String(http.StatusOK, "用户描述")
	userController.success(c)
}

func (userController UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/add.html", gin.H{
		"now": models.GetUnix(),
		"day": models.GetDate(),
	})
}

// 上传单个文件
func (UserController UserController) UploadSingleFile(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	if err != nil {
		// http.StatusBadRequest 400
		c.String(http.StatusBadRequest, "get from err: %s", err.Error())
		return
	}
	filename := "./static/" + filepath.Base(file.Filename)
	fmt.Println(filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		// http.StatusBadRequest 400
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded successfully with field username=%s,", file.Filename, username)
}

// 上传多个文件
func (UserController UserController) UploadMultipleFiles(c *gin.Context) {
	username := c.PostForm("username")
	face1, err1 := c.FormFile("face1")
	face2, err2 := c.FormFile("face2")
	if err1 != nil {
		c.String(http.StatusBadRequest, "get from err: %s", err1.Error())
		return
	}
	if err2 != nil {
		c.String(http.StatusBadRequest, "get from err: %s", err2.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "文件上传成功",
		"username": username,
	})
}
