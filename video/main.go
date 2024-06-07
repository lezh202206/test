package main

import (
	"fmt"
	request2 "test/video/request"
	"test/video/videer"
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

	option := videer.Options{
		Playlist:         false,
		Items:            "",
		ItemStart:        1,
		ItemEnd:          0,
		ThreadNumber:     10,
		Cookie:           "",
		EpisodeTitleOnly: false,
	}

	data, err := bili.Action(url, html, option)
	if err != nil {
		return
	}
	fmt.Println(data)
}
