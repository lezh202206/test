package main

import (
	"fmt"
	"time"
)

func sendData(ch chan map[int32]int32) {
	data := map[int32]int32{2: 4}
	ch <- data // 向通道发送数据
}

func receiveData(ch chan map[int32]int32) {
	msg := <-ch // 从通道接收数据
	if v, ok := msg[2]; ok {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan map[int32]int32) // 创建一个通道

	go sendData(ch)    // 启动一个 Goroutine 发送数据
	go receiveData(ch) // 启动一个 Goroutine 接收数据

	time.Sleep(time.Second) // 主 Goroutine 等待一段时间
}
