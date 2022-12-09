package main

import (
	"gin/common/function"
	"gin/core"
	"gin/core/redis"
	"gin/routers"
	"log"

	"github.com/fvbock/endless"
)

func main() {
	//注册路由
	function.ShowLogo()
	core.InitDb()
	redis.InitRedis()
	r := routers.SetupRouter()

	if err := endless.ListenAndServe(":8800", r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")
}
