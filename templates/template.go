package templates

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"go/format"
	"os"
	"strings"
	"text/template"
)

// 模板输出;
type templates struct {
	text    string      //模板的内容
	source  interface{} //模板源数据源
	funcMap template.FuncMap
}

func Templates(text string, source interface{}, funcMap template.FuncMap) *templates {
	return &templates{text, source, funcMap}
}

// 模板解析
func (t *templates) Execute() ([]byte, error) {
	// 替换占位符，将其转行为反引号
	t.text = strings.Replace(t.text, "${backquote}", "`", -1)
	tmpl, err := template.New("model").
		Funcs(t.funcMap).
		Parse(t.text)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	writer := bufio.NewWriter(buffer)
	err = tmpl.Execute(writer, t.source)
	writer.Flush()
	return buffer.Bytes(), err
}

// 模板输出
func (t *templates) Write(filename string, isFormat bool) error {
	// 模板输出;
	bytes, err := t.Execute()
	if err != nil {
		return errors.Wrap(err, "解析模板错误")
	}
	if isFormat {
		//格式化文件
		newBytes, err := format.Source(bytes)
		if err != nil {
			fmt.Println(errors.Wrap(err, "格式化模板错误").Error())
		} else {
			bytes = newBytes
		}
	}
	// 写入文件;
	if err = os.WriteFile(filename, bytes, 0660); err != nil {
		return err
	}
	return nil
}
