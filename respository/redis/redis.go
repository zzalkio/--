package redis

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// 初始化连接

// Init 初始化连接
func Init(cfg *conf.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":6379",
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("redis连接成功")
	return nil
}

func Close() {
	_ = client.Close()
}
