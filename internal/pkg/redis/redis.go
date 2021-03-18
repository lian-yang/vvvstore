package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var rdb *redis.Client

// redis配置
type Config struct {
	Host     string
	Port     int
	Password string
	Db       int
}

// 初始化redis客户端
func InitRedisClient() error {
	var conf Config
	if err := viper.UnmarshalKey("redis", &conf); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.Db,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return err
	}

	return nil
}

// 获取实例
func GetRedisClient() *redis.Client {
	return rdb
}