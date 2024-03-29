package tracing

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTrace_startSpan(t *testing.T) {
	NewTrace("test")
	s1 := StartSpan(context.Background(), "test1")
	s1.sourceOpenTracing.SetBaggageItem("k", "k1")
	k, v := s1.Inject()
	fmt.Println(k, v)
	s1.SetTags(map[string]interface{}{"k1": "111"})
	s1.End()
	time.Sleep(time.Second * 3)
}
