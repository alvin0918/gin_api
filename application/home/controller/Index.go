package HomeController

import (
	HomeModel "github.com/alvin0918/gin_api/application/home/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Nav(c *gin.Context) {

	var (
		res   map[int]map[string]string
		count string
		err   error
	)

	if res, count, err = HomeModel.GetNav(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "OK",
		"status":  false,
		"results": res,
		"count":   count,
	})
	return
}

func Banner(c *gin.Context) {
	var (
		res   map[int]map[string]string
		count string
		err   error
	)

	if res, count, err = HomeModel.GetBanner(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "OK",
		"status":  false,
		"results": res,
		"count":   count,
	})
	return
}

func Footer(c *gin.Context) {

	var (
		res   map[int]map[string]string
		count string
		err   error
	)

	if res, count, err = HomeModel.GetFooter(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "OK",
		"status":  false,
		"results": res,
		"count":   count,
	})
	return
}
