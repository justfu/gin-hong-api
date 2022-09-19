package entity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义 Result 结构体
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
}

// 定义错误码
const (
	// 成功
	CODE_SUCCESS int = 200

	//失败
	CODE_ERROR int = 10000

	//自定义...
)

// 设置返回数据
func SetSuccess(c *gin.Context, data interface{}) {
	res := Result{}
	res.Data = data
	res.Code = CODE_SUCCESS
	res.Message = "ok"
	res.Time = time.Now().Unix()
	c.JSON(http.StatusOK, res)
}
