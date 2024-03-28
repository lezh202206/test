package t_viper

import (
	"sync"
	"time"
)

type JwtBase struct {
	Secret        string `yaml:"secret"`
	Expire        int32  `yaml:"expire"`
	RefreshExpire int32  `yaml:"refresh_expire"`
	SecretBytes   []byte
	// access_token过期时间，时间：分
	ExpireMin time.Duration
	// refresh_access_token过期时间，时间：分
	RefreshExpireMin time.Duration
}

var (
	jwtOnce sync.Once
	jwtData JwtBase
)

func JWTViper() JwtBase {
	jwtOnce.Do(func() {
		Viperr("jwt.yaml", &jwtData)
	})
	// 转换对象，方便使用
	jwtData.SecretBytes = []byte(jwtData.Secret)
	jwtData.ExpireMin = time.Duration(jwtData.Expire) * time.Minute
	jwtData.RefreshExpireMin = time.Duration(jwtData.RefreshExpire) * time.Minute
	return jwtData
}
