package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

func TestTemplates(t *testing.T) {
	input := strReplace("Hello, [[.name]]! Your age is [[.age]] and your sex is [[.sex]].")
	jsonStr := `{"sex": "男", "age": 180, "name": ["title2", "title1"]}`

	// 定义一个 map 用于存储解析后的数据
	data := make(map[string]interface{})
	// 解析 JSON 字符串
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}
	//data = convSlice(data)

	// 创建模板
	tmpl := template.New("tmpl")

	// 解析模板字符串
	_tmpl, err := tmpl.Parse(input)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}

	// 应用模板到数据
	var buf bytes.Buffer
	err = _tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return
	}

	// 获取替换后的字符串
	result := buf.String()
	fmt.Println(result)

}

func strReplace(input string) string {
	input = strings.ReplaceAll(input, "[[", "{{")
	input = strings.ReplaceAll(input, "]]", "}}")
	return input
}
