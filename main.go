package main

import (
	"gin/common/env"
	"gin/common/function"
	"gin/routers"
	"log"

	"github.com/fvbock/endless"
)

func main() {
	//注册路由
	function.ShowLogo()
	r := routers.SetupRouter()
	env.Active().Value()

	if err := endless.ListenAndServe(":8800", r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")
}
