package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf16"
)

const DefaultSheetName = "Sheet1"

// ExcelWriter 为excel导出方法, 通过设置表头, 指定列名,列宽,指定默认空值,列下拉项等属性
// 当前导出仅支持 xlsx文件(不再支持xls)
type excelWriter struct {
	file   *excelize.File
	sheets []*Sheet
}

func NewExcelWriter() Writer {
	return &excelWriter{}
}

func (e excelWriter) Write(fileName string, sheets []*Sheet) ([]byte, error) {
	e.setTypeByField(sheets)

	file := excelize.NewFile()
	defer file.Close()
	for _, sheet := range sheets {
		file.NewSheet(sheet.Name)
		sw, err := file.NewStreamWriter(sheet.Name)
		if err != nil {
			return nil, err
		}

		styleID, err := file.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "777777"}})
		if err != nil {
			return nil, err
		}

		// 写表头
		if err := e.writeHeader(sw, sheet, styleID); err != nil {
			return nil, err
		}

		// 下拉项
		if err := e.writeDropList(file, sheet.Name, sheet.Headers); err != nil {
			return nil, err
		}

		rowIndex := 2
		// 行号从第2行开始(第一行为表头)
		for i := 0; i < sheet.itemsValue.Len(); i++ {
			if err := e.writeRow(sw, sheet, rowIndex, sheet.itemsValue.Index(i)); err != nil {
				return nil, err
			}
			rowIndex++
		}

		if err := sw.Flush(); err != nil {
			return nil, err
		}
	}

	file.DeleteSheet("Sheet1")
	file.SaveAs(fmt.Sprintf("/Users/lezh/Desktop/GO/src/test/%s.xlsx", fileName))

	buffer, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (e excelWriter) writeHeader(streamWriter *excelize.StreamWriter, sheet *Sheet, styleID int) (err error) {
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return err
	}

	var headerRow = make([]interface{}, 0, len(sheet.Headers))
	for _, v := range sheet.Headers {
		headerRow = append(headerRow, v.Name)
	}

	err = streamWriter.SetRow(cell, headerRow, excelize.RowOpts{StyleID: styleID, Height: 20, Hidden: false})
	if err != nil {
		return err
	}
	return nil
}

func (e excelWriter) writeRow(streamWriter *excelize.StreamWriter, sheet *Sheet, rowIndex int, item reflect.Value) (err error) {
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	row := make([]interface{}, 0, len(sheet.Headers))
	for _, header := range sheet.Headers {
		var field reflect.Value
		field = item.FieldByName(header.Field)
		// 表示找到了列
		if field.IsValid() {
			val := field.Interface()
			if field.Kind() == reflect.Slice {
				// 字段类型是切片, 需要将其组合转换为字符串(逗号隔开)
				joinVal := ""
				for i := 0; i < field.Len(); i++ {
					joinVal += fmt.Sprintf("%v, ", field.Index(i).Interface())
				}
				if joinVal != "" {
					// 移除最后的分号
					joinVal = strings.Trim(joinVal, ", ")
				}

				val = joinVal
			}

			// 显示空文本的字段
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
				// 其他类型当decimal处理
				row = append(row, val)
			default:
				row = append(row, val)
			}
		}
	}

	axis, err := excelize.CoordinatesToCellName(1, rowIndex)
	if err != nil {
		return err
	}
	if err := streamWriter.SetRow(axis, row); err != nil {
		return err
	}

	return nil
}

func (e excelWriter) setTypeByField(sheets []*Sheet) {
	for _, sheet := range sheets {
		var itemsValue = reflect.ValueOf(sheet.Items)
		if itemsValue.Kind() == reflect.Ptr {
			itemsValue = itemsValue.Elem()
		}
		if itemsValue.Kind() != reflect.Slice {
			panic("参数items数据项类型必须为slice")
		}
		sheet.itemsValue = itemsValue
	}

}

func (e excelWriter) writeDropList(f *excelize.File, sheetName string, headers []*Header) error {
	isCreateSheet := false
	dropSheetName := sheetName + "sheetName"
	for colIndex, header := range headers {
		// excel的列索引从1开始
		colIndex = colIndex + 1

		dropLen := len(header.DropList)
		if dropLen <= 0 {
			continue
		}

		colName, err := excelize.ColumnNumberToName(colIndex)
		if err != nil {
			return err
		}
		// 要有下拉的框
		dvRange := excelize.NewDataValidation(true)

		formula := strings.Join(header.DropList, ",")
		if excelize.MaxFieldLength > len(utf16.Encode([]rune(formula))) {
			dvRange.Sqref = fmt.Sprintf("%s%d:%s%d", colName, 2, colName, 1000)
			err = dvRange.SetDropList(header.DropList)
			if err != nil {
				return err
			}
		} else {
			if !isCreateSheet {
				isCreateSheet = true
				f.NewSheet(dropSheetName)
			}
			// 将下拉写入列数据
			for i, s := range header.DropList {
				// 第二行开始写入正式数据
				if err := f.SetCellStr(dropSheetName, colName+strconv.Itoa(i+2), s); err != nil {
					return err
				}
			}
			// excel验证数据的起始位置
			sqrefAxis1, _ := excelize.CoordinatesToCellName(colIndex, 2, true)
			sqrefAxis2, _ := excelize.CoordinatesToCellName(colIndex, dropLen+1, true)
			dvRange.Sqref = fmt.Sprintf("%s!%s:%s", dropSheetName, sqrefAxis1, sqrefAxis2)
			dvRange.SetSqrefDropList(dvRange.Sqref)
		}

		err = f.AddDataValidation(sheetName, dvRange)
		if err != nil {
			return err
		}
	}
	if isCreateSheet {
		// 隐藏该工作表格
		if err := f.SetSheetVisible(dropSheetName, false); err != nil {
			return err
		}
	}
	return nil
}
