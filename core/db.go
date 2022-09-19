package core

import (
	"fmt"
	"gin/config"
	"gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	host := config.Config.DB.Host
	database := config.Config.DB.Database
	user := config.Config.DB.User
	port := config.Config.DB.Port
	password := config.Config.DB.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//使用mysql 连接池
	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(60)
	sqlDB.SetConnMaxLifetime(60)
	sqlDB.SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}
	return db
}

func Db() *gorm.DB {
	return dbConnect()
}

func CreateTable() {
	//创建bs表
	Db().Model(&model.BsConfig{}).Debug().AutoMigrate(&model.BsConfig{})
	Db().Model(&model.XueqiuUgc{}).Debug().AutoMigrate(&model.XueqiuUgc{})
	Db().Model(&model.XueqiuUserSelect{}).Debug().AutoMigrate(&model.XueqiuUserSelect{})
	Db().Model(&model.XueqiuUserSelectCombination{}).Debug().AutoMigrate(&model.XueqiuUserSelectCombination{})
	Db().Model(&model.BsUserSelectCombinationHistory{}).Debug().AutoMigrate(&model.BsUserSelectCombinationHistory{})
}
