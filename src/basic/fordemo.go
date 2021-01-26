package main

import "fmt"

func main() {
	/*
		for循环：某些代码会被多次执行
		语法：
		1：标准写法
		for 表达式1；表达式2；表达式3 {
			循环体
		}
		2：同时省略表达式1和表达式3
		for 表达式2 {
		}
		相当于 while (条件)
		3:同时省略三个条件
		for {
		}
		相当于while(true)
		注意点：在for循环中，省略了表达式2，就相当于直接作用域true上
	*/
	// 主要是i作用于有变化
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

}
