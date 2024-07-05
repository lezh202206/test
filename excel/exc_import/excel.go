package exc_import

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

// 读取的工作表
type ReaderSheet struct {
	// 列定义
	Columns []*Column
	// 数据项,必须为struct的切片
	Items interface{}
	// 实体字段名称(Column.Field)与excel列的隐射
	ColMap map[string]int32
	// 读取的sheet的表头
	_headers []string
	// 内部对象, 当前excel表格读取的数据行(使用完必须关闭)
	_excelRows *excelize.Rows
}

// 工作表
type Column struct {
	// 标题
	Title string
	// 字段名
	Field     string
	Index     int32
	Format    tFunc
	_excelCol int
	Header    string
	// 对应匹配excel表头名称,必填
	Matches []string
}

type tFunc func(any2 string) (string, error)

func DateFormat(val string, layout string) (string, error) {
	// 如果取到的值是小数（表示是时间戳）
	v, _ := DoubleParse(val)
	if v > 0 {
		date, err := excelize.ExcelDateToTime(v, false)
		if err != nil {
			return "", err
		}
		return date.Format(layout), nil
	} else {
		return val, nil
	}
}

func DoubleParse(val string) (float64, error) {
	if val == "" {
		return 0, nil
	}

	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Int64Parse(val string) (int64, error) {
	if val == "" {
		return 0, nil
	}

	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
