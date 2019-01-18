package main

import (
	"os"
	"fmt"
)

//模拟linux echo 命令
//for循环是go语言中唯一的循环方式
//for init condition post {
//init(初始) condition(条件) post(后置) 都可以省略，省略之后则是无穷循环

func main()  {
	var s, sep string
	for i := 1; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}