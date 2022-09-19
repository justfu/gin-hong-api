package handler

import (
	"fmt"
	"gin/common/alarm"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

// 自定义异常中间件 支持报警提示
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				msgType := errorToString(r)
				c.JSON(http.StatusBadGateway, gin.H{
					"code": "10000",
					"msg":  msgType,
					"data": nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()

	}
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		//error错误类型 发送钉钉发
		msg := v.Error()
		e := printStackTrace()
		alarm.DingDing(e)
		return msg
	default:
		return r.(string)
	}
}

// 打印堆栈信息
func printStackTrace() string {
	var str string
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		str += fmt.Sprintf("\n > %s:%d (0x%x)", file, line, pc)
	}
	return str
}
