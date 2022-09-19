package model

import (
	"gin/config"
)

// BsUserSelectCombinationHistory  组合调仓历史' COLLATE = 'utf8mb4_general_ci
type BsUserSelectCombinationHistory struct {
	ID                 int     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid                int     `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`
	RbHid              int     `gorm:"column:rb_hid" db:"rb_hid" json:"rb_hid" form:"rb_hid"`
	RebalancingId      int     `gorm:"column:rebalancing_id" db:"rebalancing_id" json:"rebalancing_id" form:"rebalancing_id"`
	TargetVolume       float64 `gorm:"column:target_volume" db:"target_volume" json:"target_volume" form:"target_volume"`
	StockName          string  `gorm:"column:stock_name" db:"stock_name" json:"stock_name" form:"stock_name"`
	StockSymbol        string  `gorm:"column:stock_symbol" db:"stock_symbol" json:"stock_symbol" form:"stock_symbol"`
	Price              float64 `gorm:"column:price" db:"price" json:"price" form:"price"`
	NetValue           float64 `gorm:"column:net_value" db:"net_value" json:"net_value" form:"net_value"`
	Weight             float64 `gorm:"column:weight" db:"weight" json:"weight" form:"weight"`
	TargetWeight       float64 `gorm:"column:target_weight" db:"target_weight" json:"target_weight" form:"target_weight"`
	PrevWeight         float64 `gorm:"column:prev_weight" db:"prev_weight" json:"prev_weight" form:"prev_weight"`
	PrevTargetWeight   float64 `gorm:"column:prev_target_weight" db:"prev_target_weight" json:"prev_target_weight" form:"prev_target_weight"`
	PrevWeightAdjusted float64 `gorm:"column:prev_weight_adjusted" db:"prev_weight_adjusted" json:"prev_weight_adjusted" form:"prev_weight_adjusted"`
	PrevVolume         float64 `gorm:"column:prev_volume" db:"prev_volume" json:"prev_volume" form:"prev_volume"`
	PrevPrice          float64 `gorm:"column:prev_price" db:"prev_price" json:"prev_price" form:"prev_price"`
	PrevNetValue       float64 `gorm:"column:prev_net_value" db:"prev_net_value" json:"prev_net_value" form:"prev_net_value"`
	Proactive          int     `gorm:"column:proactive" db:"proactive" json:"proactive" form:"proactive"`
	PrevTargetVolume   float64 `gorm:"column:prev_target_volume" db:"prev_target_volume" json:"prev_target_volume" form:"prev_target_volume"`
	Addtime            string  `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
	Updatetime         string  `gorm:"column:updatetime" db:"updatetime" json:"updatetime" form:"updatetime"`
}

func (BsUserSelectCombinationHistory) TableName() string {
	return config.Config.DB.Prefix + "user_select_combination_history"
}
