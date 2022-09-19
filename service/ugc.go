package service

import (
	"gin/core"
	"gin/model"
)

func GetUserUgcByUid(uid string) []model.XueqiuUgc {
	var xueqiuUgcModel []model.XueqiuUgc
	core.Db().Where("uid = ?", uid).Order("id desc").Find(&xueqiuUgcModel)
	return xueqiuUgcModel
}
