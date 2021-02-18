package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本数据类型转换
	*/
	//1。bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1) // str -> bool
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1) // bool,true

	ss1 := strconv.FormatBool(b1)   // bool -> str
	fmt.Printf("%T,%s\n", ss1, ss1) // string,true

	// 2.整数
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64) // str -> int(10)
	i3, err := strconv.ParseInt(s2, 2, 64)  // str -> int(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2) // int64,100
	fmt.Printf("%T,%d\n", i3, i3) // int64,4 也就是100转化成10进制为4

	ss2 := strconv.FormatInt(i2, 10)
	ss3 := strconv.FormatInt(i3, 2)
	fmt.Printf("%T,%s\n", ss2, ss2) // string,100
	fmt.Printf("%T,%s\n", ss3, ss3) // string,100

	//itoa(),atoi()
	i4, err := strconv.Atoi("-42") // str->int
	fmt.Printf("%T,%d\n", i4, i4)  // int,-42
	ss4 := strconv.Itoa(-42)
	fmt.Printf("%T,%s\n", ss4, ss4) // string,-42
}
