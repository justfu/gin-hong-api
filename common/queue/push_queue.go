// Package queue 推送到简易消息队列/*
package queue

import (
	"encoding/json"
	"gin/core/redis"
	"time"
)

type queue struct {
	Topic    string      `json:"topic"`
	Data     interface{} `json:"data"`
	PushTime int64       `json:"pushTime"`
}

func PushQueue(topic string, data interface{}) {
	queue := queue{
		Topic:    topic,
		Data:     data,
		PushTime: time.Now().Unix(),
	}
	json, err := json.Marshal(queue)

	if err != nil {
		panic("json格式不正确")
	}

	redis.Lpush("queue_topic", string(json))

}
