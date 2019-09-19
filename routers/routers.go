package routers

import (
	UserController "github.com/alvin0918/gin_api/application/user/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine)  {

	var (
		model *gin.RouterGroup
	)

	// User模块
	model = r.Group("/user")
	{
		model.GET("/", UserController.Index)
		model.POST("/userLogin", UserController.UserLogin)
	}
}