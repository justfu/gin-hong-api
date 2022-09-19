package model

import (
	"gin/config"
	"time"
)

// BsConfig  网页配置表' COLLATE = 'utf8_general_ci
type BsConfig struct {
	ID      int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name    string    `gorm:"column:name;type:varchar(255)" db:"name" json:"name" form:"name"` //  配置描述' collate 'utf8_general_ci
	Key     string    `gorm:"column:key;type:varchar(255)" db:"key" json:"key" form:"key"`
	Value   string    `gorm:"column:value;type:varchar(255)" db:"value" json:"value" form:"value"`
	Addtime time.Time `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
	Update  time.Time `gorm:"column:update" db:"update" json:"update" form:"update"`
}

func (BsConfig) TableName() string {
	return config.Config.DB.Prefix + "config"
}
