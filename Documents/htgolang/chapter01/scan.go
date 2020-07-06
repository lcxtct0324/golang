package main

import "fmt"

func main() {

	name := ""
	fmt.Println("请输入你的名字:")

	fmt.Scan(&name) //使用指针指向name变量的的地址,实现传参

	fmt.Println("你输入的名字是: ", name)

	age := 0
	fmt.Println("请输入你的年龄:")

	fmt.Scan(&age) //使用指针指向age变量的的地址,实现传参

	fmt.Println("你输入的年龄是: ", age)

}
