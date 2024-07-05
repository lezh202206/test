package exc_import

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"log"
	"reflect"
	"strings"
)

// excelReader 为excel读取方法
type excelReader struct {
	file   *excelize.File
	sheets []*ReaderSheet
	// 限制读取行数, 默认-1不限制
	limit int

	// 是否严格校验错误，例如数据类型转换或者数据格式化错误
	// 如果开启严格校验，当excel读取的内容和字段类型不一致或者格式化出错，会直接抛错
	// 默认为true, 开启严格校验
	strictValidate bool
	// 是否读取原始的excel的值
	// 如果不开启, 当显示单元格设置了格式化,会读取到格式化的值
	// 默认为true
	rawCellValue bool
}

type ReaderOption struct {
}

func NewExcelReader() *excelReader {
	reader := &excelReader{
		// -1表示不限制数量
		limit: -1,
	}
	return reader
}

// ExcImport 导入
func (excl *excelReader) ExcImport(filePath string, data []*ReaderSheet) (err error) {
	excl.sheets = data

	err = excl.validate()
	if err != nil {
		return err
	}

	// 读取excel文件
	excl.file, err = excelize.OpenFile(filePath)
	if err != nil {
		panic("文件读取失败")
	}
	defer func(file *excelize.File) {
		err := file.Close()
		if err != nil {

		}
	}(excl.file)
	return excl.read()
}

func (excl *excelReader) validate() (err error) {
	// 校验header
	for i, sheet := range excl.sheets {
		for j, col := range sheet.Columns {

			if col.Field == "" {
				return errors.Errorf("Sheet-%v第%v列的field不能为空", i, j)
			}
			if col.Title == "" {
				return errors.Errorf("Sheet-%v列%v的Title不能为空", i, col.Field)
			}

			col.Matches = append(col.Matches, col.Title)
			// 初始化没有匹配任何列
			col._excelCol = -1
		}

		// 校验必须传入数据对象
		if sheet.Items == nil {
			return errors.Errorf("Sheet-%v的Items不能为nil", i)
		}
	}

	return nil
}

func (excl *excelReader) read() (err error) {
	for i, sheet := range excl.sheets {
		sheetName := excl.file.GetSheetName(i)
		if sheetName == "" {
			continue
		}
		// 获取所有行
		sheet._excelRows, err = excl.file.Rows(sheetName)
		if err != nil {
			log.Fatal(err)
		}
		defer sheet._excelRows.Close()

		err = excl.matchHeader(sheet)
		if err != nil {
			panic("表头 匹配失败")
		}

		if excl.limit != 0 {
			err = excl.setRows(sheet)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// matchHeader 匹配表头
func (excl *excelReader) matchHeader(sheet *ReaderSheet) (err error) {
	// 如果还有下一行
	if sheet._excelRows.Next() {
		// 取该行 的所有数据
		sheet._headers, err = sheet._excelRows.Columns()
		if err != nil {
			return err
		}
	}
	// 映射 到结构体中 通过  Excel的 Title <=>  结构体的 Title
	sheet.ColMap = make(map[string]int32)
	for index, title := range sheet._headers {
		for _, item := range sheet.Columns {
			if strings.Contains(item.Title, title) {
				item._excelCol = index
				item.Header = title
				sheet.ColMap[item.Field] = int32(index)
			}
		}
	}
	return nil
}

// setRows 匹配表头
func (excl *excelReader) setRows(sheet *ReaderSheet) (err error) {
	var dataRowNum = 0
	sliceVal := reflect.ValueOf(sheet.Items)
	for sheet._excelRows.Next() {
		dataRowNum++
		if excl.limit > 0 && dataRowNum > excl.limit {
			break
		}
		// 注意，取值的时候需要指定为RawCellValue:true
		row, err := sheet._excelRows.Columns(excelize.Options{RawCellValue: excl.rawCellValue})
		if err != nil {
			return err
		}
		// 实例化对象, sliceType.Elem()获取切片的元素
		destType := sliceVal.Type().Elem()
		// 传入的目标对象是否是指针
		isdestPtr := false
		if destType.Kind() == reflect.Ptr {
			destType = destType.Elem()
			isdestPtr = true
		}
		// 必须是结构体
		if destType.Kind() != reflect.Struct {
			return errors.New("Items对象的元素必须为结构体")
		}

		val := reflect.New(destType)

		if err = excl.setRow(sheet, val.Elem(), row); err != nil {
			return err
		}
		if isdestPtr {
			sliceVal = reflect.Append(sliceVal, val)
		} else {
			sliceVal = reflect.Append(sliceVal, val.Elem())
		}
	}
	sheet.Items = sliceVal.Interface()
	return nil
}

// 设置单行数据
func (excel *excelReader) setRow(sheet *ReaderSheet, val reflect.Value, row []string) error {
	rowLen := len(row)

	// 遍历列赋值
	for _, col := range sheet.Columns {
		field := val.FieldByName(col.Field)
		if !field.IsValid() {
			return errors.Errorf(fmt.Sprintf("实体中不存在字段%v, 请检查导出列字段的定义", col.Field))
		}

		// 防止读取的行列数与表头列数不一样, 需要判断长度
		if col._excelCol >= 0 && col._excelCol < rowLen {
			rowVal := row[col._excelCol]
			// 去除前后空格
			rowVal = strings.TrimSpace(rowVal)
			// 先格式化
			if col.Format != nil {
				newVal, err := col.Format(rowVal)
				if err != nil && excel.strictValidate {
					return err
				}
				rowVal = newVal
			}

			switch field.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				parseVal, err := Int64Parse(rowVal)
				if err != nil && excel.strictValidate {
					return err
				}
				field.SetInt(parseVal)
			case reflect.Float32, reflect.Float64:
				parseVal, err := DoubleParse(rowVal)
				if err != nil && excel.strictValidate {
					return err
				}
				field.SetFloat(parseVal)
			default:
				field.SetString(rowVal)
			}
		}
	}

	return nil
}
