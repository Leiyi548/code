package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
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
	userController.failed(c)
}
