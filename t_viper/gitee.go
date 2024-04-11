package t_viper

import "sync"

type GiteeConfig struct {
	Owner  string `yaml:"owner"` // 仓库所属空间地址(企业、组织或个人的地址path)
	Repo   string `yaml:"repo"`  // 仓库路径(path)
	Secret string `yaml:"secret"`
}

var (
	giteeOnce   sync.Once
	giteeConfig GiteeConfig
)

func GiteeViper() GiteeConfig {
	giteeOnce.Do(func() {
		Viperr("gitee.yaml", &giteeConfig)
	})
	return giteeConfig
}
