package model

import (
	"gin/config"
)

type BsLog struct {
	ID      int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Content string `gorm:"column:content" db:"content" json:"content" form:"content"`
	Addtime string `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
}

func (BsLog) TableName() string {
	return config.Config.DB.Prefix + "log"

}
