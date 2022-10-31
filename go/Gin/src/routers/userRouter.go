package routers

import (
	"ginStudy/src/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/user")
	{
		adminRouters.GET("/", user.UserController{}.Index)
		adminRouters.GET("/userList", user.UserController{}.List)
		adminRouters.GET("/description", user.UserController{}.Description)
		adminRouters.GET("/add", user.UserController{}.Add)
		adminRouters.POST("/UploadMultipleFiles", user.UserController{}.UploadMultipleFiles)
		adminRouters.POST("/UploadSingleFile", user.UserController{}.UploadSingleFile)
	}
}
