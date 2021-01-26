package main

import (
	"fmt"
)

func main() {
	/*
		if语句其他写法
		if 初始化语句;条件{

		}
	*/
	// 初始化变量只能在if语句内部访问，外部会报异常
	if num := 4; num > 0 {
		fmt.Println("整数...", num)
	}
	//fmt.Println(num)  undefined: num
}
