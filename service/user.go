package service

import (
	"fmt"
	"gin/config"
	"gin/core"
	"gin/core/redis"
	"gin/extend"
	"gin/model/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Admin struct {
	id       uint
	username string
	password string
}

// 获取TOKEN
func GetToken() {
	j := &extend.JWT{SigningKey: []byte(config.Config.Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        uuid.UUID{},
		ID:          100,
		NickName:    "DSADSADSA",
		Username:    "DSADSADSADS",
		AuthorityId: 111,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		panic("获取token失败!")
		return
	}
	fmt.Println(token)
}

func CalUserInfo() []map[int]string {
	map1 := make([]map[int]string, 0)
	for i := 0; i < 1000; i++ {
		map1 = append(map1, map[int]string{1: "sdadsa"})
	}
	return map1
}

func GetUserInfoById(id int64) map[string]interface{} {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var result = map[string]interface{}{}
	db.Table("admin").Where("id = " + strconv.FormatInt(id, 10)).Find(&result)
	return result
}

func GetUserInfoById2(id int64) map[string]interface{} {
	var result = map[string]interface{}{}
	core.Db().Table("admin").Where("id = " + strconv.FormatInt(id, 10)).Find(&result)
	//log.Println(core.Db)
	fmt.Println(result)
	return result
}

func GetRedisKey(key string) interface{} {
	res, err := redis.Get(key)
	if err != false {
		return res
	} else {
		return ""
	}
}

func Rpush(key string) interface{} {
	_ = redis.Rpush("DSADSA", "DSADSADSA")
	return nil
}
