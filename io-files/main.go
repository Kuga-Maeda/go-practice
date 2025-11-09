package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. ファイルに書き込む
	content := "こんにちは！これはGoで書かれたテキストです。\n"
	err := os.WriteFile("sample.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("書き込みエラー:", err)
		return
	}

	fmt.Println("✔️　sample.txtに書き込み完了")

	// 2. ファイルを読み込む
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("読み込みエラー:", err)
		return
	}

	fmt.Println("📕　読み込んだ内容:")
	fmt.Println(string(data))
}
