package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
	"sync"
	"test/t_viper"
)

const TraceId = "tracing-id"

var (
	once sync.Once
)

func NewTrace(serviceName string) {
	once.Do(func() {
		// 设置上报网关(后续可以改为多模式单选：入文件记录、agent代理记录)
		sender := transport.NewHTTPTransport(t_viper.TraceViper().Addr)
		//sender := transport.NewHTTPTransport("t_viper.TraceViper().EsServer")
		// 初始化 链路系统：1.服务名、2.上报规则、3.记录器
		tracer, _ := jaeger.NewTracer(
			serviceName,
			// 设置采样方案(const表示每个span都记录)
			jaeger.NewConstSampler(true),
			jaeger.NewRemoteReporter(sender, jaeger.ReporterOptions.Logger(jaeger.StdLogger)),
			jaeger.TracerOptions.CustomHeaderKeys(&jaeger.HeadersConfig{TraceContextHeaderName: TraceId}),
		)
		// 设置 tracer到全局
		opentracing.SetGlobalTracer(tracer)
	})
}

func (s *Span) Inject() (key, value string) {
	if s.sourceOpenTracing != nil {
		temp := map[string]string{}
		s.sourceOpenTracing.Tracer().Inject(s.sourceOpenTracing.Context(), opentracing.TextMap, opentracing.TextMapCarrier(temp))
		for k, v := range temp {
			return k, v
		}
	}
	return "", ""
}

func StartSpan(ctx context.Context, operationName string, opts ...SpanOption) *Span {
	var temp []opentracing.StartSpanOption

	for _, opt := range opts {
		if t := opt.parse(); t != nil {
			temp = append(temp, t.(opentracing.StartSpanOption))
		}
	}

	s, ctx := opentracing.StartSpanFromContext(ctx, operationName, temp...)
	return &Span{
		ctx:               ctx,
		sourceOpenTracing: s,
	}
}
