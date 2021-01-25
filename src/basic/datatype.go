package main

/*
内容：数据类型
*/

import "fmt"

func main() {
	/*
		Go语言数据类型：
		1.基本数据类型
			布尔类型：bool
				取值：true，false
			数值类型：
				整数
					int8: -128-127
					int16: -32768-32767
					int32
					int64
					unint8: 0-255
					byte: uint8
					rune: int32
				浮点：生活中的小数 默认保留小数点后6位
				复数
			字符串：string
		2.符合数据类型
			array，slice，map，function，pointer，struct，interface，channel
	*/

	// 1.布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T,%t\n", b1, b1)
	b2 := false
	fmt.Printf("%T,%t\n", b2, b2)

	//2.整数
	var i1 int8 // -128-127
	i1 = 100
	fmt.Println(i1)
	//i1 = 200  constant 200 overflows int8

	// 3.byte rune
	var i2 uint8
	i2 = 100
	var i3 byte
	i3 = i2
	fmt.Println(i2, i3)

	var i4 int32
	i4 = 200
	var i5 rune
	i5 = i4
	fmt.Println(i4, i5)

	// 4.浮点 默认保留小数点后6位 类型推断：float64 取决于系统
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.67
	fmt.Printf("%T,%.2f\n", f1, f1)
	fmt.Printf("%T,%f\n", f2, f2)
	/*
		float32,3.14
		float64,4.670000
	*/

}
