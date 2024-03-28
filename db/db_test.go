package db

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDbConf(t *testing.T) {
	type Customer struct {
		Id        int32
		Telephone string
		Name      string
	}
	u := &Customer{}

	err := New(context.TODO()).Write.Where("merchant_id = ? AND telephone = ?", 240, "15879832530").Find(u).Error

	time.Sleep(time.Second * 3)
	fmt.Println(u, err)
}
