package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (baseController BaseController) success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (baseController BaseController) failed(c *gin.Context) {
	c.String(http.StatusBadRequest, "failed")
}
