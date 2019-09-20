package routers

import (
	HomeController "github.com/alvin0918/gin_api/application/home/controller"
	UserController "github.com/alvin0918/gin_api/application/user/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	var (
		home *gin.RouterGroup
		user *gin.RouterGroup
	)

	// User模块
	user = r.Group("/user")
	{
		user.GET("/", UserController.Index)
		user.POST("/userLogin", UserController.UserLogin)
	}

	// User模块
	home = r.Group("/home")
	{
		home.GET("/nav/header/", HomeController.Nav)
		home.GET("/nav/footer/", HomeController.Footer)
		home.GET("/banners/", HomeController.Banner)
	}
}
