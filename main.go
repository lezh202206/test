package main

import (
	"fmt"
	"os/exec"
)

func downloadBilibiliVideo(videoURL string) error {
	cmd := exec.Command("annie", "-o", "%(title)s.%(ext)s", videoURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("下载视频失败：%v\n%s", err, output)
	}
	fmt.Println("视频下载完成！")
	return nil
}

func main() {
	videoURL := "https://www.bilibili.com/video/BV1Qz42117gt/?spm_id_from=333.1007.tianma.1-1-1.click&vd_source=7866b2f011975afa8886d726f57673d2" // 替换为你想下载的 Bilibili 视频链接
	err := downloadBilibiliVideo(videoURL)
	if err != nil {
		fmt.Println(err)
	}
}
