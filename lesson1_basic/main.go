package main

import "fmt"

func main() {
	// 名前、年齢、好きな言語
	name := "前田"
	age := 35
	favLang := "Java"

	// 文章として出力
	fmt.Println("名前", name)
	fmt.Println("年齢", age)
	fmt.Println("好きな言語", favLang)

	// 繰り返し（for文）で3回挨拶
	for i := 1; i <= 3; i++ {
		fmt.Println("こんにちは、Go学習中です！ 回数:", i)
	}

	// 条件分岐（if文）
	if age < 30 {
		fmt.Println("まだ若いですね！")
	} else {
		fmt.Println("経験が豊富ですね！")
	}

	// selfIntro関数を呼び出してみる
	selfIntro(name, age, favLang)
}

func selfIntro(name string, age int, favLang string) {
	fmt.Println("自己紹介関数の呼び出し↓")
	fmt.Println("名前", name)
	fmt.Println("年齢", age)
	fmt.Println("好きな言語", favLang)
}
