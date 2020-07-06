package main

import "fmt"

var version string = "1.0" //函数外定义变量，可以不调用变量

func main() {
	// 定义一个string类型的变量me
	/*
		变量名需要满足标识符命名规则
		1. 必须有非空的unicode字符串组成，数字 ，_
		2. 不能以数字开头
		3. 不能为go的关键字
		4. 避免和go预定义标识符冲突 例如：true/false
		5. 驼峰式
		6. 标识符区分大小写
		7. 函数内定义变量，需调用变量才可以。
	*/

	var me string

	me = "seonghoon ra"
	fmt.Println(me)

	var name, user string = "kk", "dog"
	fmt.Println(name, user)

	var (
		age    int     = 30
		height float64 = 1.68
	)

	fmt.Println(age, height)

	var (
		s = "kk"
		a = 30
	)
	fmt.Println(s, a)

	var ss, aa = "kk", 30
	fmt.Println(ss, aa)

	//这是一个简短声明,只能在函数内部使用
	isBoy := false
	fmt.Println(isBoy)

	ss, aa, isBoy = "silene", 33, true
	fmt.Println(ss, aa, isBoy)

	fmt.Println(s, ss)
	s, ss = ss, s
	fmt.Println(s, ss)
}
