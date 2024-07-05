package exc_import

import (
	"testing"
	"time"
)

func TestImport(t *testing.T) {
	type Item struct {
		Name     string
		Code     string
		Serial   int32
		Birthday string
		Amount   string
	}
	headers := []*Column{
		{
			Title: "姓名",
			Field: "Name",
		},
		{
			Title: "条码",
			Field: "Code",
		},
		{
			Title: "顺序号",
			Field: "Serial",
		},
		{
			Title: "生日",
			Field: "Birthday",
			Format: func(val string) (string, error) {
				return DateFormat(val, time.DateOnly)
			},
		},
	}
	sheets := []*ReaderSheet{{
		Columns: headers,
		// 测试传入的数据对象是指针的切片
		Items: []*Item{},
	},
	}

	err := NewExcelReader().ExcImport("/Users/lezh/Downloads/会员档案240624153155.xlsx", sheets)
	if err != nil {
		panic("")
	}
	t.Log(sheets[0].Items)
}
