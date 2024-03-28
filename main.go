package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// 定义模板
	tmpl := `
User Information:
Name: {{ .Name }}
Email: {{ .Email }}
`

	// 创建一个 User 对象
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// 解析模板
	t := template.Must(template.New("user").Parse(tmpl))

	// 应用模板并输出到标准输出
	err := t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
