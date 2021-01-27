package main

import "fmt"

func main() {
	/*
		goto语句：无条件地转移到过程中指定的行
	*/
	var a = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a的值为：%d\n", a)
		a++
	}

	fmt.Println("==============")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere // 错误跳转
			}
		}
		// 手动返回，避免执行进入标签
		return
	}
breakHere:
	fmt.Println("done...")

	fmt.Println("统一错误处理")
	/*
			err := firstCheckError()
			if err != nil {
				go onExit
			}

			err := secondCheckError()
			if err != nil {
				go onExit
			}
			fmt.Println("done")
			return

		onExit:
				fmt.Println(err)
				exitProcess()
	*/
}