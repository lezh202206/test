package main

import (
	"context"
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

	headerList := []*excel.Header{
		{Name: "商户ID", Field: "MerchantId", ColWidth: 8},
		{Name: "OpenID", Field: "Openid", ColWidth: 14},
	}
	sheets := []*excel.Sheet{{
		Name:    "test_1",
		Headers: headerList,
		Items:   &exportList,
	}}

	_, err := excel.NewExcelWriter().Write("xxx", sheets)
	if err != nil {
		return
	}
}
