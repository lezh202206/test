package redis

import (
	"context"
	"testing"
	"time"
)

func Test_redis_set(t *testing.T) {
	New(context.Background()).Set("saasdasdd", "asdas123123123", time.Second*10)
}

func Test_redis_HmSet(t *testing.T) {
	New(context.Background()).HmSet("test_HM", "title", "标题1")
}

func Test_redis_HmGet(t *testing.T) {
	data, err := New(context.Background()).HmGet("test_HM", "title")
	if err != nil {
		return
	}
	t.Log(data)
}
