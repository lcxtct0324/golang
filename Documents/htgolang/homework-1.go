package main

import "fmt"

func main() {
	num1, num2, num3 := 1, 1, 1

	for num1 <= 9 {
		if num2 <= num1 {
			num3 = num1 * num2
			fmt.Print("   ", num1, "*", num2, "=", num3, "   ")
			num2++
		} else {
			num2 = 1
			num1++
			fmt.Print("\n")
		}
	}

}
