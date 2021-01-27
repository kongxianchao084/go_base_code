package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		生成随机数random
			伪随机数，根据一定的算法公式算出来的
		math/rand
	*/
	// 生成0-9随机数  种子数没变，所以随机数不会变化
	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}
	// 改变种子数，生成的随机会变化  int64类型种子
	rand.Seed(1001)
	num2 := rand.Intn(10)
	fmt.Println(num2)

	// 要想随机数一直变化，使用时间去作为种子数即可
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1) // time.Time
	// 时间戳：距离1970年1月1日之间时间差值 秒
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1) // 1611757086
	// 纳秒
	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2) // 1611757125055809000

	// 1:真正生成随机数
	// step1：设置变化的种子数 设置时间戳
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		//step2:调用生成随机数的函数
		fmt.Println("--->", rand.Intn(100))
	}
	// 2:获取指定范围随机数
	/*
		[15,76]
		[3,48]
		Intn(n)  [0-n)
	*/
	num3 := rand.Intn(46) + 3 // 获取3-48间随机数
	fmt.Println(num3)
	num4 := rand.Intn(62) + 15 // 获取15-76间随机数
	fmt.Println(num4)
}
