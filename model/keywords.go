package model

// BsKeywords  关键字' COLLATE = 'utf8mb4_general_ci
type BsKeywords struct {
	ID      int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name    string `gorm:"column:name" db:"name" json:"name" form:"name"` //  关键词名称' collate 'utf8mb4_general_ci
	Flag    int64  `gorm:"column:flag" db:"flag" json:"flag" form:"flag"` //  1 有效 0 失效
	Addtime string `gorm:"column:addtime" db:"addtime" json:"addtime" form:"addtime"`
}

func (BsKeywords) TableName() string {
	return "bs_keywords"
}
