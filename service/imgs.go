package service

import (
	"fmt"
	"gin/common/function"
	"gin/config"
	"gin/core/redis"
	"github.com/bangbaoshi/wordcloud"
	"github.com/syyongx/php2go"
	"image/color"
	"log"
)

func GetImg() string {
	key := php2go.Date("20060102", php2go.Time()-60*60*24)

	res := redis.ZREVRANGE(key, 0, 50)

	if len(res) == 0 {
		log.Println("暂时没有词云生成")
	}

	////需要写入的文本数组
	textList := res
	//文本角度数组
	angles := []int{0, 15, -15, 90}
	//文本颜色数组
	colors := []*color.RGBA{
		&color.RGBA{0x0, 0x60, 0x30, 0xff},
		&color.RGBA{0x60, 0x0, 0x0, 0xff},
		&color.RGBA{0x73, 0x73, 0x0, 0xff},
	}

	yesterday := php2go.Date("20060102", php2go.Time()-60*60*24)

	imgName := fmt.Sprintf("%s.png", yesterday)

	imgPath := config.GetProjectTruePath() + "/imgs/" + imgName

	fmt.Println(config.GetProjectTruePath() + "/fonts/xin_shi_gu_yin.ttf")
	fmt.Println(config.GetProjectTruePath() + "/imgs/water.png")
	fmt.Println(imgPath)
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	//设置对应的字体路径，和输出路径
	render := wordcloud.NewWordCloudRender(60, 8,
		config.GetProjectTruePath()+"/fonts/xin_shi_gu_yin.ttf",
		config.GetProjectTruePath()+"/imgs/water.png", textList, angles, colors, imgPath)
	//开始渲染
	render.Render()

	fmt.Println(2323232323)

	url := config.Config.Domain.Domain + "/files/" + imgName
	//url = "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2ed3a4d1-c56b-4bdd-8158-af3690f2b2b7/2a7a95d1-76e4-409a-9139-622786f25e9f.jpg"

	title := "今日词云推送"
	content := fmt.Sprintf("## 今日词云 \n > ![词云](%s)", url)

	fmt.Println(content)
	//发送钉钉消息
	function.SendToDingDing(config.Config.DINGding.Ciyunpush, title, content)
	return url
}
