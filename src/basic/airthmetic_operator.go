package main

import "fmt"

func main() {
	/*
		算数运算符：+、-、*、/、%、==、--
	*/
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	c := 10
	d := 3
	div := c / d
	remain := c % d
	fmt.Printf("%d / %d = %d\n", c, d, div)
	fmt.Printf("%d %% %d = %d\n", c, d, remain)

	e := 3
	e++
	fmt.Println(e) // 4

	e--
	fmt.Println(e) // 3
}
