package main

import "fmt"

func main() {
	total := 0
	index := 0
	max := 100

START:
	index += 1
	total += index
	if index == max {
		goto END
	}
	goto START
END:
	fmt.Println(total)
}
