package main

import "fmt"

// Person という構造体を定義
type Person struct {
	Name string
	Age  int
}

func main() {
	// 構造体のインスタンス作成
	p1 := Person{Name: "前田", Age: 22}
	p2 := Person{Name: "山田", Age: 25}

	// 構造体のフィールドにアクセスして出力
	fmt.Println("名前:", p1.Name, "年齢:", p1.Age)
	fmt.Println("名前:", p2.Name, "年齢:", p2.Age)
}
