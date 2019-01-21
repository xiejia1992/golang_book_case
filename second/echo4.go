package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "是否换行")
var sep = flag.String("s", " ", "连接符号")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}


