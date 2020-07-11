package main

import "fmt"

func main() {
	// 零值 nil
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	//赋值
	//使用现有变量 取地址 &name
	age := 32

	age2 := age

	fmt.Printf("%T, %#v\n", &age, &age)
	fmt.Printf("%T, %#v\n", &age2, &age2)

	pointerInt = &age

	fmt.Println(pointerInt)
	fmt.Println(*pointerInt)

	*pointerInt = 325

	fmt.Println(age, age2) //即使修改 age的值，并不会更改age2的值 因为指针地址不同
	fmt.Printf("%T, %#v\n", &age, &age)

	pointerString = new(string)
	fmt.Printf("%#v, %#v\n", pointerString, *pointerString)
}
