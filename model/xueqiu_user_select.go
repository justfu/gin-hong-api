package model

import "gin/config"

type XueqiuUserSelect struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Symbol   string `json:"content"`
	Name     string `json:"xueqiuuid"`
	Type     int    `json:"type"`
	Remark   string `json:"remark"`
	Exchange string `json:"exchange"`
	Created  int    `json:"created"`
	Category int    `json:"category"`
	Addtime  string `json:"addtime"`
}

func (XueqiuUserSelect) TableName() string {
	return config.Config.DB.Prefix + "xueqiu_user_select"
}
