package main

import "fmt"

// Person構造体
type Person struct {
	Name string
	Age  int
}

// PersonにGreetメソッドを追加
func (p Person) Greet() {
	fmt.Println("こんにちは、私の名前は", p.Name, "です。年齢は", p.Age)
}

// Ageを増やすメソッド（ポインタレシーバ）
func (p *Person) Birthday() {
	p.Age++
}

func main() {
	p := Person{Name: "前田", Age: 22}

	// メソッド呼び出し
	p.Greet()
	p.Birthday()
	p.Greet()
}
