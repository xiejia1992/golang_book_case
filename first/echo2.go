package main

import (
	"os"
	"fmt"
)

//与上一个版本的区别在于初始化s, sep赋值方式不同
//for 循环方式也不同
//如果只想看结果，一句代码就能实现：fmt.Println(os.Args[1:])

func main()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}