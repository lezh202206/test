package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	topFoods := []string{
		"一心一味",
		"黄大妈",
		"珍香鹅",
	}

	lowFoods := []string{
		"猪脚饭",
		"潮汕自助",
		"一碗享",
	}

	// 获取筛选后的元素
	selectedElement := selectElement(time.Now(), topFoods, lowFoods)

	// 打印结果
	fmt.Printf("Selected element: %d\n", selectedElement)

	getWeather()
}

// 根据给定的日期和数组，返回一个筛选后的元素
func selectElement(date time.Time, array1, array2 []string) string {
	day := date.Day()
	rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	fmt.Println(randomNumber)

	// 计算当月的第15号之后的一周的日期范围
	after15th := 15
	before22nd := 22

	// 根据日期选择优先数组
	if day >= after15th && day < before22nd {
		return array1[randomNumber%len(array1)]
	}
	// 其他时间，优先选择array2
	return array2[randomNumber%len(array2)]
}

const apiKey = "9c04cc71e38da872caee6c887eb23bbd" // 替换为你的 API 密钥
const city = "shenzhen"                           // 替换为你想查询的城市
const apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather() {
	url := fmt.Sprintf(apiURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get weather data: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
