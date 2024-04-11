package enum_gen

import (
	"os"
	"path"
	"strings"
	"test/gen/json/define"
	"test/gen/json/sources"
	"test/gen/json/tmpl"
	"test/templates"
)

type enumsGen struct {
	root   string
	source []*define.EnumsSource
}

const defaultPath = "/Users/lezh/Desktop/GO/src/test"

func EnumsGen() *enumsGen {
	return &enumsGen{
		root: path.Join(defaultPath, "gen"),
	}
}

func (eg enumsGen) Gen(fileName string) {
	eg.source = sources.GetSources(path.Join(eg.root, fileName))

	eg.genFolder(fileName)

	eg.genFile(fileName)
}

// 首先， 生成各级目录
// 目录为 enums/文件名称(去掉文件格式)/定义的枚举
func (e *enumsGen) genFolder(fileName string) {
	folder := strings.Split(fileName, ".")[0]
	if folder == "" {
		panic("未指定文件名称.")
	}
	// 每个文件会单独创建一个枚举根目录
	servPath := path.Join(e.root, folder)
	// 首先，删除该目录下所有文件
	if err := os.RemoveAll(servPath); err != nil {
		panic(err)
	}
	// 创建模块目录;
	if err := os.MkdirAll(servPath, 0766); err != nil {
		panic(err)
	}
	// 重新组目录
	for _, val := range e.source {
		if err := os.Mkdir(path.Join(servPath, val.Group), 0766); err != nil {
			panic(err)
		}
	}
}

// 生成文件
func (e *enumsGen) genFile(fileName string) error {
	for _, group := range e.source {
		// enums目录
		path := path.Join(e.root, strings.Split(fileName, ".")[0], group.Group, "e.go")
		if err := templates.Templates(tmpl.Enums, group, nil).Write(path, true); err != nil {
			return err
		}
	}

	return nil
}
