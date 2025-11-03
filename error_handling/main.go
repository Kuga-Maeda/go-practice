package main

import (
	"errors"
	"fmt"
)

// divideは a を b で割り、エラーがあれば返す
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("０で割ることはできません")
	}
	return a / b, nil
}

func main() {
	// 正常な計算
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("エラー", err)
	} else {
		fmt.Println("結果:", result)
	}

	// エラーになる計算
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー", err)
	} else {
		fmt.Println("結果:", result)
	}

	// panic/recover の例
	safeDivide(10, 0)
}

// panic/recover　を使った安全な割り算
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from panic:", r)
		}
	}()

	if b == 0 {
		panic("0で割ることはできません")
	}

	fmt.Println("結果:", a/b)
}
