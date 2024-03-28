package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// CustomParameters 结构体表示自定义参数
type CustomParameters struct {
	Param1 string `json:"param1"`
	Param2 int    `json:"param2"`
	// 添加其他参数...
}

func main() {
	r := gin.Default()

	// 定义一个处理程序，该处理程序将处理POST请求
	r.POST("/send-request", func(c *gin.Context) {
		// 解析POST请求中的JSON数据
		var requestBody CustomParameters
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
			return
		}

		// 替换为您要发送请求的URL
		apiURL := "https://example.com/api"

		// 发送POST请求
		responseBody, err := sendPostRequest(apiURL, requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}

		// 返回第三方服务的响应
		c.String(http.StatusOK, string(responseBody))
	})

	// 启动Gin服务器
	r.Run(":8080")
}

// sendPostRequest 发送POST请求并返回响应
func sendPostRequest(url string, requestBody CustomParameters) ([]byte, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
