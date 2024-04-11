package t_viper

import (
	"strings"
	"sync"
)

type ServerConfig struct {
	SrvName string
	Host    string
	Port    string
	Addr    string
}

type httpCnf struct {
	Mode       string                  `yaml:"mode"`
	Service    map[string]string       `yaml:"service"`
	ServConfig map[string]ServerConfig `yaml:"-"`
	SecretKey  map[string]string       `yaml:"secret_key"`
}

var (
	httpOnce sync.Once
	httpCof  httpCnf
)

func HttpViper() httpCnf {
	httpOnce.Do(func() {
		Viperr("http.yaml", &httpCof)
		httpConfByService(&httpCof)
	})
	return httpCof
}

func httpConfByService(servConfig *httpCnf) {
	httpCof.ServConfig = make(map[string]ServerConfig)
	for k, v := range servConfig.Service {
		splits := strings.Split(v, ":")
		if len(splits) != 2 {
			panic("配置文件错误")
		}

		c := ServerConfig{
			SrvName: k,
			Host:    splits[0],
			Port:    splits[1],
			Addr:    v,
		}
		httpCof.ServConfig[k] = c
	}
	return
}
