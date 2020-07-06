package main

import "fmt"

func main() {

	msg := "我的名\\n字是Seonghoon"
	fmt.Println(msg)
	//赋值 + +=
	msg += "---GOGO" //追加赋值
	fmt.Println(msg)

	//索引 切片 ascii
	msg = "abcdef"
	fmt.Printf("%T %#v %c\n", msg[0], msg[0], msg[0])
	fmt.Println(msg[1:3])

	//len 字节大小
	fmt.Println(len(msg))
}
