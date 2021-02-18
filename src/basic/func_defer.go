package main

import (
	"fmt"
)

func main() { // 外围函数
	/*
		defer的词义："延迟"
		在go语言中，使用defer关键字来延迟一个函数或方法的执行

		1。defer函数或方法：一个函数或方法的执行被延迟了
		2。defer的方法：
			A：对象.close()，临时文件的删除
			B：go语言中关于异常的处理，使用panic()和recover()
				panic()函数用于引发恐慌，导致程序终中断执行
				recover()函数用于恢复程序的执行，recover()语法上要求必须在defer中执行
		3。如果多个defer函数：先延迟的后执行，后延迟的先执行
		4。defer函数传递参数的时候:defer函数调用时，值已经传递给参数，只是暂时不执行函数中的代码而已
		5。defer函数注意点:
			当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
			当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正结束
			当外围函数中的代码引发引起恐慌时，只有其中所有延迟函数都执行完毕后，该运行恐慌才会真正被扩展至调用函数
	*/
	defer func1("hello")
	fmt.Println("12345")
	defer func1("world")
	fmt.Println("王二狗")
	/*
		defer函数会在外围函数执行完成后，在执行defer语句
		defer语句"栈结构"："hello"先放--栈底，再放"world"--栈顶
		因此"world"先输出，在输出"hello"
		12345
		王二狗
		world
		hello
	*/
	a := 2
	fmt.Println(a) // 2
	defer func2(a) // 2
	a++
	fmt.Println("main:", a) // 3
	/*
		12345
		王二狗
		2
		main: 3
		func2()函数中打印a： 2
		world
		hello
	*/
	fmt.Println(func3())
	/*
		12345
		王二狗
		2
		main: 3
		func3()函数的执行。。。
		haha
		hehe
		0
		func2()函数中打印a： 2
		world
		hello
	*/
}

func func1(s string) {
	fmt.Println(s)
}

func func2(a int) {
	fmt.Println("func2()函数中打印a：", a) // func2()函数中打印a： 2
}

func func3() int {
	fmt.Println("func3()函数的执行。。。")
	defer func1("hehe")
	defer func1("haha")
	return 0
}
