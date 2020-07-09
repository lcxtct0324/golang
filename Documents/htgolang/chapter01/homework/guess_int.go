package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	/*
		猜数字游戏 生成随机整数[0, 100) 提示用户再控制台输入猜测的数字 比较，当用户输入较大，提示太大了
		当用户输入太小，提示太小了 当用户输入正确，提示经过N次对了，太聪明了 用户最多猜5次，如果5次内都没有猜正确，
		提示太笨了，游戏结束
		扩展: 当成功或失败后，提示用户是否继续，输入：yes, y, Y则继续，重新生成随机数，让用户猜测

		游戏逻辑：
		1.进入游戏
		2.输入字符，转换类型并存储
		3.判断准确度，提示大小
		4.yes重新开始游戏，其他退出

		代码逻辑：
		for死循环，通过isQuit变量的值，控制是否开始游戏
		进入游戏后，通过for死循环进行单词游戏，通过usedGameNum计数控制最大尝试次数，
		通过userStr获取用户输入并判断是否符合数字要求，符合计算游戏次数，不符合要求用户重新输入。
		符合数字规则时，进行游戏逻辑判断。
	*/
	rand.Seed(time.Now().Unix()) //给rand函数设置一个随机种子
	userStr := ""                //初始化一个错误的值，让逻辑进入重新输入环节。也用来防止用户输入非数字
	isQuit := false              //是否退出游戏
	maxGameNum := 6              //每轮最多猜几次!

	for {
		if isQuit { //退出游戏
			fmt.Println("GoodBye!")
			break
		} else { //进入游戏
			usedGameNum := 1
			gameNum := rand.Intn(100)
			fmt.Println(gameNum)
			fmt.Println("[0]:欢迎来带猜数字游戏!")
			for {
				if usedGameNum < maxGameNum {
					if inNum, err := strconv.Atoi(userStr); err == nil {
						if inNum == gameNum {
							fmt.Printf("[1]:%d次就猜对了,你太聪明了!\n", usedGameNum)
							fmt.Print("[2]:在玩一局？ yes/no: ")
							fmt.Scan(&userStr)
							if userStr != "y" && userStr != "Y" && userStr != "yes" {
								isQuit = true
							}
							break
						} else if inNum > gameNum {
							fmt.Println("数字过大")

						} else if inNum < gameNum {
							fmt.Println("数字过小")
						}
						userStr = ""
						usedGameNum++

					} else {
						fmt.Println("请输入一个整数: ")
						fmt.Scan(&userStr)
					}
				} else {
					fmt.Printf("[1]:%d次都错了,你太笨了! 你没有资格重新开始了,游戏结束! \n", usedGameNum)
					isQuit = true
					break
				}
			}
		}
	}
}
