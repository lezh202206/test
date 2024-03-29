package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

type SpanOption interface {
	parse() interface{}
}

type Span struct {
	ctx               context.Context
	sourceOpenTracing opentracing.Span
}

type TraceData map[string]interface{}

// SetTags 设置标签描述
func (s *Span) SetTags(data TraceData) *Span {
	if s.sourceOpenTracing != nil {
		for k, v := range data {
			s.sourceOpenTracing.SetTag(k, v)
		}
	}
	return s
}

// End 链路结束
func (s *Span) End() {
	if s.sourceOpenTracing != nil {
		s.sourceOpenTracing.Finish()
	}
}
