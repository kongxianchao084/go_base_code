package main

import "fmt"

func main() {
	/*
		函数的返回值
			函数返回的结果，必须和函数定义的一致：类型，个数，顺序
			可以使用_来舍弃多余的返回值
			如果一个函数定义了有返回值，那么函数中有分支、循环，那么要保证，无论执行了那个分支，都要有return语句执行
	*/
	res := getSum()
	fmt.Println(res)
	res1 := getSum2()
	fmt.Println(res1)
	// 多个返回值
	res2, res3 := rectangle(5, 3)
	fmt.Printf("周长：%.2f，面积：%.2f\n", res2, res3)
	res4, res5 := rectangle2(5, 3)
	fmt.Printf("周长：%.2f，面积：%.2f\n", res4, res5)

}

func getSum() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func getSum2() (sum int) { // 定义函数时，指定要返回的数据是哪一个
	//上面已经指定输出为sum，已经初始化，此处不可在初始化sum变量
	//sum := 0  no new variables on left side of :=
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

// 多个返回值：计算长方形周长和面积
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (peri float64, area float64) { // 初始化
	peri = (len + wid) * 2
	area = len * wid
	return
}
