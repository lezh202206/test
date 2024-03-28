package excel

import (
	"github.com/xuri/excelize/v2"
	"reflect"
)

// 表头
type Header struct {
	// 表头名称
	Name string
	// 对应的字段(数据项的字段,用于赋值)
	Field string
	// 列宽, 如果不只设置则使用默认宽度
	// 默认宽度值为8, 大概英文|数字:8个,中文4个
	ColWidth float64
	// 指定空数据, 值为这些数据的时候, excel表格显示空(用于清理一些0值问题), 注意, 必须设置的数据类型与字段类型一致才会起效
	EmptyValue interface{}
	// 下拉选项, 仅用于导出表头方法
	DropList []string
	// 设置列是否必填,仅用于导出表头的时候, 标注表头为红色
	Required bool
}

// 工作表
type Sheet struct {
	// sheet名称, 名称不允许为Sheet1(该工作表为默认表, 会被删除)
	Name string
	// 表头
	Headers []*Header
	// 数据项,必须为struct的切片
	Items interface{}

	// 内部使用的对象(数据源对象)
	itemsValue reflect.Value
	//获取流式写入器
	streamWriter *excelize.StreamWriter
}

type Writer interface {
	// 写入数据(使用流式写入)
	Write(fileName string, sheets []*Sheet) ([]byte, error)
}
