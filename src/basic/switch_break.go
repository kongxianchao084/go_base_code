package main

import "fmt"

func main() {
	/*
		switch中的break和fallthrough语句
		break：可以使用在break中，也可以使用在for循环中
		fallthrough：穿透下一个case语句，无需判断直接输出；必须位于最后一行
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		break // 强制结束switch语句
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
	}
	fmt.Println("main...over...")
	/*
		我是熊二
		main...over...
	*/

	fmt.Println("=============")
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
		//fmt.Println("fallthrough必须位于最后一行") fallthrough statement out of place
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
	/*
		第二季度
		第三季度
	*/
}
