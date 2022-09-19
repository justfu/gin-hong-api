package model

import (
	"gin/config"
)

type XueqiuUgc struct {
	Id        int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid       int    `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`
	Content   string `gorm:"column:content;type:varchar(255)" db:"content" json:"content" form:"content"`
	Xueqiuuid int    `gorm:"column:xueqiuuid" db:"xueqiuuid" json:"xueqiuuid" form:"xueqiuuid"`
	CreatedAt int64  `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	Addtime   string `gorm:"column:addtime;type:varchar(255)" db:"addtime" json:"addtime" form:"addtime"`
}

func (XueqiuUgc) TableName() string {
	return config.Config.DB.Prefix + "xueqiu_ugc"
}
