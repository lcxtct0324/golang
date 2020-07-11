package main

import "fmt"

func main() {
	var age8 int8 = 31
	var age = 31
	fmt.Printf("%T, %#v, %d\n", age8, age8, age8)
	fmt.Printf("%T, %#v, %d\n", age, age, age)

	// 算数运算 + _ * / % ++ --

	a, b := 2, 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	b--

	fmt.Println(a, b) // 3, 3

	// 关系运算 > < >= <= != ==

	fmt.Println(a > b)  //false
	fmt.Println(a < b)  //false
	fmt.Println(a >= b) //ture
	fmt.Println(a >= b) //ture
	fmt.Println(a == b) //ture
	fmt.Println(a != b) //false
}
