package templates

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// 定义模板字符串
const tmplStr = `"时间：{{ range .Logs -}} {{ .CreateDt }} {{ .StatusName }} {{end -}}"`

type User struct {
	Data string
}
type Logs struct {
	CreateDt   string
	Name       string
	StatusName string
	Duration   string
	Body       string
}

type ls struct {
	LogsStatusName string
	ConfigName     string
}

type Base struct {
	Logcreatedt   string
	Configname    string
	Logstatusname string
	Logduration   string
	Logbody       string
	Usermobile    string
}

func NewTemplate() {
	// 定义模板字符串
	tmplStr := xxx()

	// 定义替换数据
	logs := Base{
		Logcreatedt:   "2024 年 03 月 30 日 14:37:19",
		Configname:    "大王",
		Logstatusname: "测试人",
		Logduration:   "2024 年 03 月 30 日",
		Logbody:       "不知道写什么",
		Usermobile:    " 110",
	}

	// 创建模板
	tmpl := template.New("tmpl")

	// 解析模板字符串
	tmpl, err := tmpl.Parse(tmplStr)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	// 应用模板到数据
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, logs)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}

	// 获取替换后的字符串
	result := buf.String()
	fmt.Println(result)
}

func xxx() string {
	_tmplStr := "{\"http\":{\"method\":\"POST\",\"url\":\"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xx\",\"body\":\"{\\n    \\\"msgtype\\\": \\\"text\\\",\\n    \\\"text\\\": {\\n        \\\"content\\\": \\\"时间：[[log.create_dt]]\\\\n任务 [[config.name]]执行[[log.status_name]]了 \\\\n耗时[[log.duration]]秒\\\\n响应：[[log.body]]\\\",\\n        \\\"mentioned_mobile_list\\\": [[user.mobile]]\\n    }\\n}\",\"header\":[{\"key\":\"\",\"value\":\"\"}]}}"

	// 定义正则表达式模式
	re := regexp.MustCompile(`\[\[(.*?)\]\]`)
	// 使用正则表达式提取字段
	matches := re.FindAllStringSubmatch(_tmplStr, -1)
	// 生成最终的字符串
	for _, match := range matches {
		field := match[1]
		replacement := "{{." + strings.Title(strings.ReplaceAll(field, ".", "")) + "}}"
		// 替换时将下划线替换为空格
		replacement = strings.ReplaceAll(replacement, "_", "")
		_tmplStr = strings.ReplaceAll(_tmplStr, "[["+field+"]]", replacement)
	}

	return _tmplStr
}
