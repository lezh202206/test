package t_viper

import "sync"

type tracingConf struct {
	// 链路追踪es的地址
	Addr string `yaml:"addr"`
}

var (
	tracingOnce sync.Once
	tracingC    tracingConf
)

// tracing
func TraceViper() tracingConf {
	tracingOnce.Do(func() {
		Viperr("tracing.yaml", &tracingC)
	})
	return tracingC
}
