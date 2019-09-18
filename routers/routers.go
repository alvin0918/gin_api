package routers

import (
	"github.com/alvin0918/gin_api/application/user/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine)  {
	r.GET("/", controller.Index)
}