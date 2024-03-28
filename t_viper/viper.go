package t_viper

import (
	"encoding/json"
	"github.com/spf13/viper"
)

func Read(fileName string) {
	viper.AddConfigPath("/Users/lezh/Desktop/GO/src/test/config") // 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
	viper.AddConfigPath(".")                                      // 把当前目录加入到配置文件的搜索路径中
	viper.SetConfigType("yaml")                                   // 设置配置文件格式；
	viper.SetConfigName(fileName)                                 //  设置配置文件名
	viper.AutomaticEnv()                                          // 设置 t_viper 查找是否有跟配置文件中相匹配的环境变量，如果有，则将该环境变量的值设置为配置项的值；

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Viperr(fileName string, data interface{}) {
	Read(fileName)
	// 获取配置数据
	configData := viper.AllSettings()
	// 将配置数据转换为 JSON 格式
	jsonData, err := json.Marshal(configData)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		panic(err)
	}
}
