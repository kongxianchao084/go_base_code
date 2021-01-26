package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		输入和输出:
			fmt包：输入、输出

		输出：
			Print()  打印
			Println()  打印换行
			Printf()  格式化打印

		格式化打印占位符：
			%v，原样输出
			%T，打印类型
			%t，bool类型
			%s，字符串类型
			%f，浮点类型
			%d，10进制的整数
			%b，2进制的整数
			%o，8进制
			%x，%X，16进制
				%x：0-9 a-f
				%X：0-9 A-F
			%c，打印字符
			%p，打印地址

		输入：
			fmt.Scanln() 一行输入多个值，通过取地址赋值给变量
			fmt.Scanf()  通过格式化输出
			bufio包：操作文件读写
	*/
	a := 100
	b := 3.14
	c := true
	d := "hello world"
	e := 'A' // 单个字符
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%.3f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%d,%c\n", e, e, e)
	fmt.Println("=========================")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	/*
		int,100
		float64,3.140
		bool,true
		string,hello world
		int32,65,A
		=========================
		100
		3.14
		true
		hello world
		65
	*/
	fmt.Println("======输入==============")
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点类型：")
	fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作地址，赋值给x和y  阻塞式
	fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)
	/*
		输入：300 3.14
		输出：x的数值：300，y的数值：3.140000
	*/
	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d，y:%f\n", x, y)
	/*
		输入: 400,5.23
		输出: x:400，y:5.230000
	*/

	fmt.Println("bufio包测试，请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n') //换行表示读取完
	fmt.Println("读到的数据：", s1)
	/*
		bufio包测试，请输入一个字符串：
		hello nihao
		读到的数据： hello nihao
	*/
}
