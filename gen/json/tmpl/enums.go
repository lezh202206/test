package tmpl

// 模板内容-数据表实体隐射,go对象;
const Enums = `package {{ .Group }}

{{ range .Enums }}
// {{ .Note }}
const (
	{{range .Items -}}
	// {{ .Note }}
	{{ .Key }} =  {{ .Value }}
	{{end -}}
)

// {{ .Note }}	
var {{.Name}}Types = map[int32]string {
	{{range .Items -}}
	// {{ .Note }}
	{{ .Key }} : "{{ .Note }}",
	{{end -}}
}

// {{ .Note }}切片	
var {{.Name}}Slice = []map[int32]string {
	{{range .Items -}}
	// {{ .Note }}
	{ {{ .Key }} : "{{ .Note }}" },
	{{end -}}
}
{{end}}
`
