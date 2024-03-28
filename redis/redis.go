package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"test/t_viper"
	"time"
)

var (
	once sync.Once
	rdb  *redis.Client
)

type Redis struct {
	Ctx context.Context
	Cli *redis.Client
}

func New(ctx context.Context) *Redis {
	once.Do(func() {
		redisCof := t_viper.RedisViper()
		// 创建 Redis 客户端连接
		rdb = redis.NewClient(&redis.Options{
			Addr:         redisCof.Addr,                                     // Redis 服务器地址
			Password:     redisCof.Password,                                 // Redis 访问密码，如果没有密码则留空
			DB:           redisCof.Db,                                       // Redis 数据库索引，默认为 0
			DialTimeout:  time.Duration(redisCof.DialTimeout) * time.Second, // Redis 数据库索引，默认为 0
			PoolSize:     redisCof.PoolSize,
			MinIdleConns: redisCof.MinIdleConns,
			IdleTimeout:  time.Duration(redisCof.IdleTimeout) * time.Second,
		})
	})

	// 连接 Redis 服务器
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return &Redis{
		Cli: rdb,
		Ctx: ctx,
	}
}

func (redis Redis) Set(key string, value interface{}, expiration time.Duration) {
	err := redis.Cli.Set(redis.Ctx, key, value, expiration).Err()
	if err != nil {
		return
	}
}

func (redis Redis) HmSet(key string, data ...interface{}) {
	err := redis.Cli.HMSet(redis.Ctx, key, data).Err()
	if err != nil {
		return
	}
}

func (redis Redis) HmGet(key string, fields string) (data []interface{}, err error) {
	data, err = redis.Cli.HMGet(redis.Ctx, key, fields).Result()
	return data, nil
}
