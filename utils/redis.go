package utils

import (
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type RedisConn struct {
	RedisHost   string `json:"redisHost" yaml:"redisHost" binding:"required"`
	RedisPort   string `json:"redisPort" yaml:"redisPort" binding:"required"`
	RedisPasswd string `json:"redisPasswd" yaml:"redisPasswd" binding:"required"`
}

func LoadRedisConfig() RedisConn {
	var redisConfig RedisConn
	data, err := os.ReadFile("D:\\beeDemo\\conf\\redis.yaml")
	if err != nil {
		fmt.Printf("打开redis配置文件出粗: %s\n", err.Error())
	}
	errUnmarshal := yaml.Unmarshal(data, &redisConfig)
	if errUnmarshal != nil {
		return RedisConn{}
	}
	return redisConfig
}

type RedisResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RedisInfo struct {
	Key        string
	Value      string
	ExpireTime int
}

func SaveToRedis(redisConn RedisConn, redisInfo RedisInfo) RedisResult {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConn.RedisHost + ":" + redisConn.RedisPort,
		Password: redisConn.RedisPasswd,
		DB:       0,
	})
	var redisResult RedisResult
	err := rdb.Set(context.Background(), redisInfo.Key, redisInfo.Value, time.Second*120).Err()
	if err != nil {
		LogToFile("Error", err.Error())
		redisResult = RedisResult{Code: 0, Msg: fmt.Sprintf("存储redis数据失败，错误: %s", err.Error())}
	} else {
		redisResult = RedisResult{Code: 1, Msg: "存储redis数据成功"}
	}
	return redisResult
}

type SearchRedisResult struct {
	Code        int
	RedisResult string
}

func SearchRedis(username string) SearchRedisResult {
	redisConn := LoadRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConn.RedisHost + ":" + redisConn.RedisPort,
		Password: redisConn.RedisPasswd,
		DB:       0,
	})
	verifyCode, err := rdb.Get(context.Background(), username).Result()
	var searchRedisResult SearchRedisResult
	if err != nil {
		LogToFile("Error", fmt.Sprintf("用户%s redis查询验证码失败，错误L %s", username, err.Error()))
		searchRedisResult = SearchRedisResult{Code: 0, RedisResult: ""}
	} else {
		searchRedisResult = SearchRedisResult{Code: 1, RedisResult: verifyCode}
	}
	return searchRedisResult
}
