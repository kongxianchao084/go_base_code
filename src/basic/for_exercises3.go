package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：三位数  100 999
		每个位上的数字的立方和，刚好等于该数字本身，那么就叫水仙花数

		268
			268 / 100 = 1 百位
			268 % 10 = 8  个位

			268 --> 26 % 10 = 6
			268 --> 68 / 10 = 6
	*/
	for i := 100; i < 1000; i++ {
		bai := i / 100
		ge := i % 10
		shi := i / 10 % 10
		if math.Pow(float64(bai), 3)+math.Pow(float64(shi), 3)+math.Pow(float64(ge), 3) == float64(i) {
			fmt.Printf("%d是水仙花数\n", i)
		}
	}
}
