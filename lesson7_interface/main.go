package main

import "fmt"

// Speaker インターフェース
type Speaker interface {
	Speak()
}

// Person 構造体
type Person struct {
	Name string
}

// PersonにSpeakメソッドを実装
func (p Person) Speak() {
	fmt.Println("こんにちは、私の名前は", p.Name, "です")
}

// Dog構造体
type Dog struct {
	Name string
}

// 　DogにSpeakメソッドを実装
func (d Dog) Speak() {
	fmt.Println(d.Name, "ワンワン！")
}

func main() {
	var s1 Speaker

	s1 = Person{Name: "前田"}
	s1.Speak()

	s1 = Dog{Name: "ポチ"}
	s1.Speak()
}
