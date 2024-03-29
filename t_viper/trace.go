package t_viper

import "sync"

type tracingConf struct {
	// 链路追踪es的地址
	EsServer string `yaml:"es_server"`
}

var tracingOnce sync.Once
var tracingC tracingConf

// tracing
func TraceViper() tracingConf {
	tracingOnce.Do(func() {
		tracingC = tracingConf{}
		Viperr("tracing.yaml", &tracingC)
	})
	return tracingC
}
