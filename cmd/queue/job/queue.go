package job

import (
	"encoding/json"
	"fmt"
	"gin/common/alarm"
	"gin/common/function"
	"gin/config"
	"gin/config/extra"
	"gin/core/redis"
	"gin/lib"
	"github.com/reugn/go-quartz/quartz"
	"github.com/tidwall/gjson"
	"reflect"
	"strconv"
	"time"
)

// PrintJob implements the quartz.Job interface.
type Queue struct {
	Desc   string
	Signal chan bool
}

// Description returns the description of the PrintJob.
func (pj *Queue) Description() string {
	return pj.Desc
}

// Key returns the unique PrintJob key.
func (pj *Queue) Key() int {
	return quartz.HashCode(pj.Description())
}

// Execute is called by a Scheduler when the Trigger associated with this job fires.
func (pj *Queue) Execute() {

	for {
		queueLen := redis.Llen("queue_topic")

		fmt.Println("当前队列长度" + strconv.FormatInt(queueLen, 10))

		if queueLen == 0 {
			time.Sleep(time.Second * 3)
			continue
		}

		if config.Config.Queue.Isopenmultithreading == 1 {
			//多线程执行
			data := redis.LpopMul("queue_topic", 30)

			if len(data) > 0 {
				dataChan := make(chan string, 30)

				//塞数据
				for _, item := range data {
					dataChan <- item
				}

				//关闭数据写通道
				close(dataChan)

				exitChan := make(chan bool, 16)

				for i := 0; i < 16; i++ {
					go execute_task_mul(dataChan, exitChan)
				}

				for i := 0; i < 16; i++ {
					<-exitChan
				}
			}

		} else {
			Data := redis.Rpop("queue_topic")

			var a interface{}
			json.Unmarshal([]byte(Data), &a)
			jsonData := a.(string)

			if jsonData != "" {
				// 消费topic
				topic := gjson.Get(jsonData, "topic").String()
				pushTime := gjson.Get(jsonData, "pushTime").Int()
				data := gjson.Get(jsonData, "data").String()

				if topic == "" {
					continue
				}

				if data == "" {
					continue
				}

				execute_task(topic, pushTime, data)
			}
		}

	}

}

func execute_task(topic string, pushTime int64, data string) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := function.ErrorToString(err)
			alarm.DingDing(errMsg)
		}
	}()

	executeFunction := extra.Queue[topic]

	if _, ok := executeFunction.(lib.GinHongQueue); !ok {
		panic("接口没有实现" + topic)
	}

	dataF := reflect.ValueOf(executeFunction)

	method := dataF.MethodByName("Run")

	if method.Kind() == reflect.Func {
		params := make([]reflect.Value, 3)
		params[0] = reflect.ValueOf(topic)
		params[1] = reflect.ValueOf(pushTime)
		params[2] = reflect.ValueOf(data)
		method.Call(params)
	}
}

//多线程执行队列任务 提升处理速度
func execute_task_mul(dataChan chan string, exitChan chan bool) {
	for {

		chanData, ok := <-dataChan

		if !ok {
			break
		}

		var a interface{}
		json.Unmarshal([]byte(chanData), &a)
		jsonData := a.(string)

		if jsonData == "" {
			continue
		}

		defer func() {
			if err := recover(); err != nil {
				errMsg := function.ErrorToString(err)
				alarm.DingDing(errMsg)
			}
		}()

		topic := gjson.Get(jsonData, "topic").String()
		pushTime := gjson.Get(jsonData, "pushTime").Int()
		data := gjson.Get(jsonData, "data").String()

		fmt.Println("正在处理" + data)

		executeFunction := extra.Queue[topic]

		if _, ok := executeFunction.(lib.GinHongQueue); !ok {
			panic("接口没有实现" + topic)
		}

		dataF := reflect.ValueOf(executeFunction)

		method := dataF.MethodByName("Run")

		if method.Kind() == reflect.Func {
			params := make([]reflect.Value, 3)
			params[0] = reflect.ValueOf(topic)
			params[1] = reflect.ValueOf(pushTime)
			params[2] = reflect.ValueOf(data)
			method.Call(params)
		}

	}
	exitChan <- true
}
