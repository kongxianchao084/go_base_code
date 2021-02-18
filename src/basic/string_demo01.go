package main

import (
	"fmt"
)

func main() {
	/*
		go中的字符串是一个字节的切片
			可以通过将其内容封装在""中来创建字符串。Go中的字符串是Unicode兼容的，并且是UTF-8编码的

		字符串是一些字节的集合

		语法：""、``
	*/

	//1。定义字符串
	s1 := "hello中国"
	s2 := `hello world`
	fmt.Println(s1)
	fmt.Println(s2)

	//2.字符串长度：字节的个数 中文占三个字节
	fmt.Println(len(s1)) // 11

	// 3。获取某个字节
	fmt.Println(s1[0])        // 返回对应的字节数值  104
	fmt.Printf("%c\n", s1[0]) // h

	// 4。遍历string
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c", s2[i])
	}

	fmt.Println()

	for _, v := range s2 {
		fmt.Printf("%c", v)
	}

	fmt.Println()

	for _, v := range s1 {
		fmt.Printf("%c", v)
	}

	fmt.Println()

	// 5.字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1) // 根据一个字节切片，构建字符串
	fmt.Println(s3)      // ABCDE

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	// 6.字符串不能修改
	//s4[0] = 0   cannot assign to s4[0] (strings are immutable)
}
