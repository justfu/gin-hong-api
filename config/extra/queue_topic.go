// 定义所有的队列结果处理函数
package extra

import (
	"gin/event/addLog"
	"gin/event/exeWords"
)

var Queue = map[string]interface{}{
	"AddLog":   &addLog.AddLog{},
	"ExeWords": &exeWords.ExeWords{},
}
