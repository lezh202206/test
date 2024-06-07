package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const apiKey = "9c04cc71e38da872caee6c887eb23bbd" // 替换为你的 API 密钥
const city = "jieyang"                            // 替换为你想查询的城市
const apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=zh_cn"

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	url := fmt.Sprintf(apiURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get weather data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get valid response: %s", body)
	}

	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Printf("Current temperature in %s: %.2f°C\n", city, weatherResponse.Main.Temp)
	fmt.Printf("Weather description: %s\n", weatherResponse.Weather[0].Description)
}
