package main

import "fmt"

func main() {
	const NAME string = "seonghoon"

	//省略类型
	//定义多个常亮 (类型相同)
	//定义多个常亮 (类型不相同)
	//定义多个常亮 省略类型
	fmt.Println(NAME)
	const AGE, HEIGHT = 30, "qier"

	fmt.Println(AGE, HEIGHT)

	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12
	)
	fmt.Println(C7, C8, C9, C10, C11, C12)

	// 枚举，const+iota
	const (
		E1 int = iota
		E2
		E3
	)

	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E1, E2, E3)
	fmt.Println(E4, E5, E6)
}
