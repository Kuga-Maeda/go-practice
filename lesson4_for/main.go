package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("回数:", i)
	}

	i := 1
	for i <= 5 {
		fmt.Println("while的な書き方:", i)
		i++
	}
}
