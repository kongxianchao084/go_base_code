package main

import "fmt"

/*
内容：变量的内存分析和注意事项
*/

// 5：简短定义方式，不能定义全局变量
// date := "星期一"  syntax error: non-declaration statement outside function body
var date string = "星期一"

func main() {
	/*
		注意事项：
		1 变量必须先定义才可以使用
		2 go语言是静态语言，要求变量的类型和赋值的类型必须一致
		3 变量名不能冲突（同一个作用域）
		4 简短定义方式，不能定义全局变量
		5 变量的零值，也叫默认值
		6 变量定义了就要使用，否则无法通过编译
	*/
	var num int
	num = 100
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)
	/*
		num的数值是：100，地址是：0x140000ae008
		num的数值是：200，地址是：0x140000ae008
	*/

	// 1：变量必须先定义才可以使用
	//fmt.Println(a)  undefined: a

	// 2：变量和赋值的类型要保持一致
	//var b string
	//b = 3.14
	//fmt.Printf("b的类型是：%T，地址是：%p，数值是：%d\n",b,&b,b)
	//cannot use 3.14 (type untyped float) as type string in assignment

	// 3：变量名不能重复
	//var num int = 300
	//fmt.Println(num) num redeclared(重新声明) in this block

	// 4：简短定义方式，左边的变量名至少有一个是新的
	num, name := 300, "孔小宝"
	fmt.Println(num, name)
	/*
		300 孔小宝
	*/

	// 5：简短定义方式，不能定义全局变量
	fmt.Println(date)

	// 6：变量的默认值
	var m1 int
	var m2 string
	var m3 float64
	var m4 bool
	fmt.Println(m1, m2, m3, m4)
	/*
		0  0 false
	*/

	// 6：变量定义了就要使用，否则无法通过编译
	//n1 := 200  n1 declared but not used
}
