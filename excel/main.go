package main

import (
	"context"
	"gorm.io/gorm/utils"
	"test/db"
	"test/excel/excel"
	"test/models"
)

func main() {
	var customer []*models.Customer
	if err := db.New(context.TODO()).Write.Where("telephone = ? AND merchant_id = ?", 15879832530, 240).Find(&customer).Error; err != nil {
		return
	}

	type exportItem struct {
		models.Customer
	}
	exportList := make([]*exportItem, 0, len(customer))

	for _, v := range customer {
		item := &exportItem{
			Customer: *v,
		}
		exportList = append(exportList, item)
	}

	var drop1 []string
	for i := 100; i < 200; i++ {
		drop1 = append(drop1, utils.ToString(i))
	}

	headerList := []*excel.Header{
		{Name: "商户ID", Field: "MerchantId", ColWidth: 8, DropList: []string{"模板", "模板1"}},
		{Name: "OpenID", Field: "Openid", ColWidth: 14, DropList: []string{"OpenID1", "OpenID2"}},
		{Name: "下拉选项", Field: "Drop", ColWidth: 14, DropList: drop1},
	}
	sheets := []*excel.Sheet{{
		Name:    "test_1",
		Headers: headerList,
		Items:   &exportList,
	}}

	_, err := excel.NewExcelWriter().Write("下拉选项", sheets)
	if err != nil {
		return
	}
}
