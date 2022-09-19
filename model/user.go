package model

import (
	"gin/config"
)

type BsXueqiuUser struct {
	ID              int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid             string `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`
	Username        string `gorm:"column:username" db:"username" json:"username" form:"username"`
	ForkCombination int64  `gorm:"column:fork_combination" db:"fork_combination" json:"fork_combination" form:"fork_combination"` //  是否关注组合
	IsNeedPush      int64  `gorm:"column:is_need_push" db:"is_need_push" json:"is_need_push" form:"is_need_push"`
	Flag            int64  `gorm:"column:flag" db:"flag" json:"flag" form:"flag"` //  状态
	Addtime         string `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
}

func (BsXueqiuUser) TableName() string {
	return config.Config.DB.Prefix + "xueqiu_user"
}
