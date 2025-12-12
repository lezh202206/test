package RadixSort

import (
	"fmt"
)

// 基数排序（对非负整数，升序）
func countingSort(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}
	return output
}

func radixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}
	return arr
}

type Person struct {
	Name   string
	Age    int
	Height int
	Score  int
}

// 稳定重排：keys 是排序后（有序）的 key 序列
func reorder(old []Person, keys []int) []Person {
	n := len(old)
	buckets := make(map[int][]Person, n)
	for i := 0; i < n; i++ {
		buckets[keys[i]] = append(buckets[keys[i]], old[i])
	}
	result := make([]Person, n)
	pos := 0
	for _, k := range keys {
		result[pos] = buckets[k][0]
		buckets[k] = buckets[k][1:]
		pos++
	}
	return result
}

func RadixSort() {
	people := []Person{
		{"A", 20, 180, 90},
		{"B", 18, 175, 88},
		{"C", 20, 175, 95},
		{"D", 20, 175, 90},
		{"E", 18, 165, 100},
	}

	// 1) Score 降序：转成升序 key = max - score
	maxScore := 0
	for _, p := range people {
		if p.Score > maxScore {
			maxScore = p.Score
		}
	}
	scoreKeys := make([]int, len(people))
	for i, p := range people {
		scoreKeys[i] = maxScore - p.Score
	}
	scoreKeys = radixSort(scoreKeys)
	people = reorder(people, scoreKeys)

	// 2) Height 升序
	heightKeys := make([]int, len(people))
	for i := range people {
		heightKeys[i] = people[i].Height
	}
	heightKeys = radixSort(heightKeys)
	people = reorder(people, heightKeys)

	// 3) Age 升序
	ageKeys := make([]int, len(people))
	for i := range people {
		ageKeys[i] = people[i].Age
	}
	ageKeys = radixSort(ageKeys)
	people = reorder(people, ageKeys)

	for _, p := range people {
		fmt.Printf("%+v\n", p)
	}
}
