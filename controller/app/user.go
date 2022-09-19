package app

import (
	"fmt"
	"gin/common/queue"
	commonService "gin/common/service"
	"gin/config"
	"gin/core/redis"
	"gin/entity"
	"gin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	http "net/http"
	"strconv"
	"time"
)

func UserHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.2com",
	})
}

func UserAdd(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		panic("ID不能为空")
	}
	parseInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	res := service.GetUserInfoById(parseInt)
	entity.SetSuccess(c, res)
}

func UserAdd2(c *gin.Context) {
	cacheKey := commonService.GetTestKey()
	data := redis.Rmember(cacheKey, func() interface{} {
		id := c.Query("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic("解析错误")
		}
		res := service.GetUserInfoById2(parseInt)

		return res
	}, 60)

	c.JSON(http.StatusOK, data)
}

func UserUgc(c *gin.Context) {
	now1 := time.Now()
	fmt.Println(now1.UnixMilli())
	uid := c.Query("uid")
	cacheKey := commonService.GetUserUgc(uid)
	now2 := time.Now()
	fmt.Println(now2.UnixMilli())
	data := redis.Rmember(cacheKey, func() interface{} {
		log.Println("走数据库")
		res := service.GetUserUgcByUid(uid)
		return res
	}, 60)
	now3 := time.Now()
	fmt.Println(now3.UnixMilli())
	redis.Tag(commonService.GetTestTag(), []string{cacheKey})

	now4 := time.Now()
	fmt.Println(now4.UnixMilli())
	entity.SetSuccess(c, data)
}

func UserUgcClear(c *gin.Context) {
	redis.TagClear(commonService.GetTestTag())
}

func GetRedis(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		panic("ID不能为空")
	}
	res := service.GetRedisKey(key)
	c.JSON(http.StatusOK, res)
}

func ClearTag(c *gin.Context) {
	redis.TagClear(commonService.GetTestTag())
}

func Rpush(c *gin.Context) {
	queue.PushQueue("AddLog", map[string]interface{}{
		"content": "1111",
	})
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "dsadsadsads"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestexeWord(c *gin.Context) {
	name := c.Query("name")
	queue.PushQueue("ExeWords", map[string]interface{}{
		"content": name, //2222
	})
}

func Test22(c *gin.Context) {
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
	fmt.Println(config.Config.DB.Host)
}

func GetToken(c *gin.Context) {
	service.GetToken()
}
