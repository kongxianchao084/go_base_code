package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		string包下关于字符串的函数
	*/
	s1 := "helloworld"
	// 1。字符串子串判断
	fmt.Println(strings.Contains(s1, "llo")) // true
	// 2。是否包含chars中任意一个字符，包含一个即为true
	fmt.Println(strings.ContainsAny(s1, "abcd")) // true
	// 3。统计字符串中子串出现的次数
	fmt.Println(strings.Count(s1, "l")) // 3
	// 4.以什么开头、以什么结尾
	s2 := "2019052525课堂笔记.txt"
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("201905的文件")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("txt文件")
	}
	// 5.获取下标  第一次出现的位置 未找到返回-1
	fmt.Println(strings.Index(s1, "l"))        // 2
	fmt.Println(strings.Index(s1, "llo"))      // 2
	fmt.Println(strings.Index(s1, "aaa"))      // -1
	fmt.Println(strings.IndexAny(s1, "aaad"))  // 9 查找chars中任意一个字符，出现在s中的位置
	fmt.Println(strings.LastIndex(s1, "l"))    // 8 从后往前找
	fmt.Println(strings.LastIndexAny(s1, "l")) // 8 从后往前找
	// 6。字符串拼接
	ss1 := []string{"abc", "world", "hello", "ruby"}
	s3 := strings.Join(ss1, "*")
	fmt.Println(s3) // abc*world*hello*ruby
	//7.字符串切割
	s4 := "123,456,aaaa,49588"
	ss2 := strings.Split(s4, ",")
	fmt.Println(ss2)
	fmt.Printf("%T\n", ss2) // []string  切片
	//8。重复
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)
	//9.替换
	s6 := strings.Replace(s1, "l", "*", 2)
	fmt.Println(s6) // he**oworld
	s7 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s7) // he**owor*d
	//10.大小写转换
	s8 := "hello WOrlD**123"
	fmt.Println(strings.ToLower(s8))
	fmt.Println(strings.ToUpper(s8))

}
