package addLog

import (
	"fmt"
	"gin/common/alarm"
	"gin/core"
	"gin/model"
	"github.com/syyongx/php2go"
	"github.com/tidwall/gjson"
)

type AddLog struct {
}

func (addLog AddLog) Run(topic string, pushTime int64, data string) {
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			alarm.DingDing(topic + "执行异常")
		}
	}()

	fmt.Println(pushTime)

	content := gjson.Get(data, "content")

	fmt.Println(content.Exists())

	if !content.Exists() {
		return
	}

	contentString := content.String()

	core.Db().Create(&model.BsLog{
		Content: contentString,
		Addtime: php2go.Date("2006-01-02 15:04:05", pushTime),
	})

}
