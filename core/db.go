package core

import (
	"fmt"
	"gin/config"
	"gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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

	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetConnMaxLifetime(20 * time.Second) //连接池里面的连接最大存活时长。
	sqlDB.SetMaxOpenConns(5)                   //连接池里最大空闲连接数

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
