package t_viper

import (
	"sync"
)

type RedisBase struct {
	Addr         string `yaml:"addr"`
	Db           int    `yaml:"db"`
	Password     string `yaml:"password"`
	DialTimeout  int32  `yaml:"dial_timeout"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
	IdleTimeout  int32  `yaml:"idle_timeout"`
}

var (
	RedisOnce sync.Once
	RedisData RedisBase
)

func RedisViper() RedisBase {
	RedisOnce.Do(func() {
		Viperr("redis.yaml", &RedisData)
	})
	// 转换对象，方便使用
	return RedisData
}
