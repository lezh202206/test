package t_viper

import (
	"fmt"
	"testing"
)

func TestDbConf(t *testing.T) {
	fmt.Printf("%+v", DBViper())
}

func TestJwtConf(t *testing.T) {
	fmt.Printf("%+v", JWTViper())
}

func TestRedisConf(t *testing.T) {
	fmt.Printf("%+v", RedisViper())
}

func TestTraceConf(t *testing.T) {
	fmt.Printf("%+v", TraceViper().Addr)
}

func TestHttpConf(t *testing.T) {
	fmt.Printf("%+v", HttpViper())
}

func TestGiteeConf(t *testing.T) {
	fmt.Printf("%+v", GiteeViper())
}
