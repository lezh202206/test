package GitPR

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"test/models"
	viper1 "test/t_viper"
)

type GiteePR struct {
	viper1.GiteeConfig
}

var (
	once sync.Once
	conf viper1.GiteeConfig
)

func NewGiteePR() GiteePR {
	once.Do(func() {
		conf = viper1.GiteeViper()
	})
	return GiteePR{conf}
}

func (conf GiteePR) SendGitee(title, head, base, assignees string) string {
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/pulls", conf.Owner, conf.Repo)
	// 准备请求的数据，将它们编码成 JSON 格式
	data := map[string]interface{}{
		"access_token": conf.Secret,
		"title":        title,     // PR 标题
		"head":         head,      // 合并分支
		"base":         base,      // 目标分支
		"assignees":    assignees, // 审查人员
	}
	// 转换为 JSON 格式
	payload, err := json.Marshal(data)
	if err != nil {
		panic("JSON 编码错误:")
	}

	// 发起 POST 请求
	resp, err := http.Post(url, "application/json;charset=UTF-8", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("", err)
		panic("POST 请求失败:")

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("响应失败:")
	}

	var List models.GiteeResponse
	err = json.Unmarshal(body, &List)
	if err != nil {
		panic("转义失败:")
	}
	fmt.Println(List.HtmlUrl)
	return List.HtmlUrl
}

func (conf GiteePR) SendBot(title string) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=991cbde3-6963-4adc-a25c-7a6402ab7d38")

	// 准备请求的数据，将它们编码成 JSON 格式
	// 构造消息内容
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":               title,
			"mentioned_list":        []string{"乐智辉"},
			"mentioned_mobile_list": []string{"15879832530"},
		},
	}
	// 将消息内容转换为 JSON 格式
	payloadBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

}
