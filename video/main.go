package main

import (
	"fmt"
	request2 "test/video/request"
	"test/video/videer/bili"
)

const (
	url     = "https://www.bilibili.com/bangumi/play/ep198218?spm_id_from=333.337.0.0"
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
