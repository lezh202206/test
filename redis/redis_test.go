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

func Test_redis_Producer(t *testing.T) {
	New(context.Background()).Producer()
}

func Test_redis_Consumer(t *testing.T) {
	New(context.Background()).Consumer()
}
