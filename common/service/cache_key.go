// Package service 缓存key值 统一管理
package service

import "gin/common/function"

func GetTestKey() string {
	return function.RunFuncName()
}

func GetUserUgc(uid string) string {
	return function.RunFuncName() + uid
}
