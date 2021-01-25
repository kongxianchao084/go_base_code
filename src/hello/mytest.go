package main

import "fmt"

/*
go run -n mytest.go   先编译，在link
go run -work mytest.go  查看生成的临时文件
go run命令在执行第二次的时候，如果发现导入的包没有变化，不再次编译代码包，而是直接静态链接进来。

go build
1.普通包执行，不会产生任何文件
2.main包，go build当前目录下生成可执行文件；go install/go build -o +路径，可以在$GOPATH/bin下生成二进制文件
go build -n  查看编译的过程

go install 编译并安装代码包
go install命令在内部分成了两步：1 生成结果文件（可执行文件或者.a包）2 编译的结果移到$GOPATH/pkg或者$GOPATH/bin
可执行文件：一般是go install安装带main函数的go文件产生
.a应用包：一般是go install安装不包含main函数的go文件产生的，没有函数入口，只能被调用

go get
go get命令会把当前代码包下在到GOPATH中第一个工作区的src目录下，并安装
go get -d  只执行下载动作，不安装
go get -u 增量下载
go get -f  使用-u时有效，检查代码包导入路径，如果下载安装的代码包所属的项目是从别人fork来的，需要使用该参数
go get -x github.com/go-errors/errors

go doc  查看包文档说明
127.0.0.1:9527
*/

func main() {
	fmt.Println("你好，Go！hahahsss")
}
