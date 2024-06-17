package main

import (
	"test/downloader"
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

	defaultDownloader := downloader.New(downloader.Options{
		Silent:         false,
		InfoOnly:       false,
		Stream:         "",
		AudioOnly:      false,
		Refer:          "",
		OutputPath:     "",
		OutputName:     "",
		FileNameLength: 255,
		Caption:        false,
		MultiThread:    false,
		ThreadNumber:   10,
		RetryTimes:     10,
		ChunkSizeMB:    1,
		UseAria2RPC:    false,
		Aria2Token:     "",
		Aria2Method:    "http",
		Aria2Addr:      "localhost:6800",
	})

	errors := make([]error, 0)
	for _, item := range data {
		if item.Err != nil {
			errors = append(errors, item.Err)
			continue
		}
		if err = defaultDownloader.Download(item); err != nil {
			errors = append(errors, err)
		}
	}
}
