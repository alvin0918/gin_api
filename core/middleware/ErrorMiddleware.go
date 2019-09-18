package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 异常处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 记录错误信息
		printRecover()

		// 处理请求
		c.Next()
	}
}

func printRecover() {

	var (
		res  interface{}
		data []byte
		err  error
	)

	if res = recover(); res != nil {
		if data, err = json.Marshal(err); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(data)
	}
}
