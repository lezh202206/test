package sources

import (
	"encoding/json"
	"os"
	"strings"
	"test/gen/json/define"
)

// 获取数据源
// fullPath:json文件名完整路径
func GetSources(fullPath string) []*define.EnumsSource {
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	// 反序列化数据;
	var enums []*define.EnumsItem
	if err := json.Unmarshal(bytes, &enums); err != nil {
		panic(err)
	}

	// 遍历，将包名称小写以及将字段首字母大写
	for _, v := range enums {
		// 枚举名称转换为驼峰形式
		v.Name = camel(v.Name)
		for _, enum := range v.Items {
			// 将枚举的字段转换为驼峰
			enum.Key = field(v.Name, enum.Key)
		}
	}
	// 将数据分组
	dic := make(map[string][]*define.EnumsItem)
	for _, v := range enums {
		dic[v.Group] = append(dic[v.Group], v)
	}

	var out []*define.EnumsSource
	for key, val := range dic {
		source := &define.EnumsSource{
			Group: key,
			Enums: val,
		}
		out = append(out, source)
	}

	return out
}

// 获取枚举项字段名称(为枚举名称+字段名称)例如： status Success Failed
func field(name string, key string) string {
	str := name + "_" + key
	return camel(str)
}

// 将值转换为大驼峰，移除下划线
func camel(str string) string {
	// 将字段组合后，显示为驼峰
	r := ""
	s := strings.Split(str, "_")
	for _, v := range s {
		r += Capital(v)
	}
	return r
}

// 首字母大写
func Capital(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
