package core

import (
	"fmt"
	"gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

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

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 50)

	//sqlDB.SetConnMaxIdleTime(time.Minute * 2)

	if err != nil {
		log.Println("init db fail!!!")
		panic(err)
	}
	return db
}

func InitDb() {
	log.Println("init db ing...")
	Db = dbConnect()
	log.Println("init db ok")
}
