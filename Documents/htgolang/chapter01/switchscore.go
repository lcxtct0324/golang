package main

import "fmt"

func main() {
	// 在控制台输入分数
	// 90分以上 A
	// 80分以上 B
	// 60分以上 C
	// 60分一下 D

	var score float32
	fmt.Print("请输入分数:")
	fmt.Scan(&score)
	fmt.Println("您输入的分数为:", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
