package routers

import (
	"gin/controller/app"
	"gin/handler"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(handler.Recover(), gin.Logger(), handler.Cors())

	GroupV1 := r.Group("/app")
	{
		GroupV1.Any("/ugc/UserUgc", app.UserUgc)
		GroupV1.Any("/ugc/UserUgcClear", app.UserUgcClear)
		GroupV1.Any("/member/UserAdd", app.UserAdd)
		GroupV1.Any("/test", app.Rpush)
		GroupV1.Any("/test22", app.Test22)
		GroupV1.Any("/TestexeWord", app.TestexeWord)
	}

	r.Static("/files", "./imgs/")

	r.Any("/getRankPic", app.GetImgs)
	r.Any("/upload", app.UploadFileTest)
	r.Any("/GetToken", app.GetToken)

	return r
}
