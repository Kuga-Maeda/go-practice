package main

import "fmt"

// 　2つの整数を受け取って、合計を返す関数
func add(a int, b int) int {
	return a + b
}

// 2つの文字列を受け取って、結合して返す関数
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	sum := add(3, 5)
	fmt.Println("3 + 5 =", sum)

	word := concat("Go", "Lang")
	fmt.Println("文字列の結合結果:", word)
}
