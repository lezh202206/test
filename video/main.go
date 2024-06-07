package main

import (
	"fmt"
	request2 "test/video/request"
	"test/video/videer/bili"
)

const (
	url     = "https://www.bilibili.com/video/BV1E142127b2/?spm_id_from=333.1007.tianma.1-1-1.click&vd_source=7866b2f011975afa8886d726f57673d2"
	referer = "https://www.bilibili.com"
)

func main() {
	html, err := request2.Get(url, referer, nil)
	if err != nil {
		return
	}

	pageData, err := bili.GetMultiPageData(html)
	if err != nil {
		return
	}
	fmt.Println(pageData)
}
