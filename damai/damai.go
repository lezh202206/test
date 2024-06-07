package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://www.bilibili.com/video/BV1Qz42117gt/?spm_id_from=333.1007.tianma.1-1-1.click&vd_source=7866b2f011975afa8886d726f57673d2" // 替换为你要下载的视频的 URL
	//
	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 创建保存视频的文件
	outFile, err := os.Create("《漫长的季节》11：最后的欢宴，是梦？是命？【墨菲】[0].mp4.download") // 指定保存的文件名
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	// 将 HTTP 响应的内容写入文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Video downloaded successfully!")
}
