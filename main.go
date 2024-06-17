package main

import (
	"fmt"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	// 设置进度条的总大小
	size := int64(1000)

	// 定义自定义模板
	tmpl := `{{counters .}} {{bar . "[" "=" ">>" "-" "]"}} {{speed .}} {{percent . | cyan}} {{rtime .}}`

	// 创建进度条
	bar := pb.New64(size).
		Set(pb.Bytes, true).
		SetMaxWidth(1000).
		SetTemplate(pb.ProgressBarTemplate(tmpl))

	// 开始进度条
	bar.Start()

	// 模拟工作
	for i := int64(0); i < size; i++ {
		time.Sleep(time.Millisecond * 10) // 模拟工作负载
		bar.Increment()                   // 进度条前进
	}

	// 完成进度条
	bar.Finish()

	fmt.Println("工作完成")
}
