package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 种子
	rand.Seed(time.Now().Unix())

	//[0-100]

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 100)
	}
}
