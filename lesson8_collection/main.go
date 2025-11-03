package main

import "fmt"

// スライスの合計を返す関数
func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// マップの値を全て出力する関数
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func main() {
	// スライス
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("合計:", sum(numbers))

	// マップ
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 95,
	}
	printMap(scores)
}
