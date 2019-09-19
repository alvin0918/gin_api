package UserController

import (
	userLogic "github.com/alvin0918/gin_api/application/user/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "Let's go!",
		"data":   make(map[string]string),
		"status": true,
	})
}

func UserLogin(c *gin.Context) {
	var (
		username string
		password string
		isTrue   bool
		err      error
	)

	// 获取参数
	username = c.PostForm("username")
	password = c.PostForm("password")

	// 参数判断
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    "参数错误！",
			"status": false,
		})
		return
	}

	if isTrue, err = userLogic.UserLogin(username, password); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    "系统错误！",
			"status": false,
		})
		return
	}

	if isTrue {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    "登录成功！",
			"status": true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "用户名或密码错误",
		"status": false,
	})
	return
}
