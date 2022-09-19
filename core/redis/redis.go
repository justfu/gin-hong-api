package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"gin/config"
	"github.com/go-redis/redis/v8"
	"github.com/syyongx/php2go"
	"time"
)

var Redis *redis.Client

var tag string = "tag_"

// 初始化redis连接
func Connect() {
	name := config.Config.REDIS.Host
	port := config.Config.REDIS.Port
	password := config.Config.REDIS.Password
	r := redis.NewClient(&redis.Options{
		Addr:     name + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB

		PoolSize: 10,
	})

	_, err := r.Ping(context.Background()).Result()

	if err != nil {
		panic("redis 连接失败!")
	}

	Redis = r
}

// Get 获取指定key值  如果redis没有这个key 返回"",false
func Get(key string) (interface{}, bool) {
	Connect()
	var ctx = context.Background()
	res, err := Redis.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	} else {
		var data interface{}

		err := json.Unmarshal([]byte(res), &data)

		if err != nil {
			panic(err)
		}

		return data, true
	}
}

// 设置指定key值
func Set(key string, value interface{}, expired int64) {

	Connect()

	dataType, _ := json.Marshal(value)

	dataString := string(dataType)

	var ctx = context.Background()
	_, err := Redis.Set(ctx, key, dataString, time.Duration(expired)*time.Second).Result()

	if err != nil {
		panic(err)
	}
}

// redis 操作已经简化
func Rmember(cacheKey string, callQuery func() interface{}, expired int64) interface{} {

	// 此处通过 redis 获取数据, 如果存在数据, 那么直接返回
	data, err := Get(cacheKey)

	if err != false {
		return data
	}

	// 当 redis 没有数据, 那么调用此方法修改 t,
	saveData := callQuery()

	Set(cacheKey, saveData, expired)

	return saveData
}

// 推队列 左推 返回当前队列数量
func Lpush(key string, value interface{}) int64 {

	Connect()

	dataType, _ := json.Marshal(value)

	dataString := string(dataType)

	var ctx = context.Background()

	res, err := Redis.LPush(ctx, key, dataString).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 右推
func Rpush(key string, value interface{}) int64 {

	Connect()

	dataType, _ := json.Marshal(value)

	dataString := string(dataType)

	var ctx = context.Background()

	res, err := Redis.RPush(ctx, key, dataString).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 右推
func ZADD(key string, member string, score float64) {

	Connect()

	var ctx = context.Background()

	items := &redis.Z{Score: score, Member: member}

	_, err := Redis.ZAdd(ctx, key, items).Result()

	if err != nil {
		panic(err)
	}

}

// 获取有序集合的成员数
func ZCARD(key string) bool {

	Connect()

	var ctx = context.Background()

	res, err := Redis.ZCard(ctx, key).Result()

	if err != nil {
		panic(err)
	}

	if res > 0 {
		return true
	} else {
		return false
	}

}

// 返回有序集合中指定成员的索引
func ZRANK(key string, member string) int64 {

	Connect()

	var ctx = context.Background()

	res, err := Redis.ZRank(ctx, key, member).Result()

	if err != nil {
		if err == redis.Nil {
			return -1
		}
	}

	return res

}

// 获取有序集合的成员数
func ZINCRBY(key string, incryment float64, member string) {

	Connect()

	var ctx = context.Background()

	_, err := Redis.ZIncrBy(ctx, key, incryment, member).Result()

	if err != nil {
		panic(err)
	}
}

func ZREVRANGE(key string, start int64, end int64) []string {
	Connect()

	var ctx = context.Background()

	res, err := Redis.ZRevRange(ctx, key, start, end).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 弹队列 左推 返回当前队列数量
func Lpop(key string) string {

	Connect()

	var ctx = context.Background()

	res, err := Redis.LPop(ctx, key).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 弹队列 右弹 返回当前队列数量
func Rpop(key string) string {

	Connect()

	var ctx = context.Background()

	res, err := Redis.RPop(ctx, key).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 弹队列 左弹多个值 返回当前队列数量
func LpopMul(key string, len int64) []string {

	Connect()

	var ctx = context.Background()

	res, err := Redis.LRange(ctx, key, 0, len-1).Result()
	Redis.LTrim(ctx, key, len, -1).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 弹队列 右弹 返回当前队列数量
func Llen(key string) int64 {

	Connect()

	var ctx = context.Background()

	res, err := Redis.LLen(ctx, key).Result()

	if err != nil {
		panic(err)
	}

	return res
}

// 删除缓存
func Del(key string) {

	Connect()

	var ctx = context.Background()

	_, err := Redis.Del(ctx, key).Result()

	if err != nil {
		panic(err)
	}
}

// 给缓存打标签
func Tag(tagKey string, keySlice []string) {
	//首先获取当前tagKey下面的缓存
	tagKeyEncryption := tag + php2go.Md5(tagKey)

	tagkeyRes, res := Get(tagKeyEncryption)

	fmt.Println("====")
	fmt.Println(res)
	fmt.Println(tagkeyRes)

	keyStr := php2go.Implode(",", keySlice)
	//没有tagkey
	if res == false {

		Set(tagKeyEncryption, keyStr, 0)

	} else {

		tagKeySlice := php2go.Explode(",", tagkeyRes.(string))

		for _, value := range keySlice {
			if php2go.InArray(value, tagKeySlice) {
				continue
			}
			tagKeySlice = append(tagKeySlice, value)
		}

		newKeyStr := php2go.Implode(",", tagKeySlice)

		Set(tagKeyEncryption, newKeyStr, 0)

	}

}

// 给缓存打标签
func TagClear(tagKey string) {
	//首先获取当前tagKey下面的缓存
	tagKeyEncryption := tag + php2go.Md5(tagKey)

	fmt.Println(tagKeyEncryption)

	tagkeyRes, res := Get(tagKeyEncryption)

	fmt.Println(tagkeyRes)

	if res == true {
		tagKeySlice := php2go.Explode(",", tagkeyRes.(string))

		for _, value := range tagKeySlice {
			Del(value)
		}

		Del(tagKeyEncryption)

	}
}
