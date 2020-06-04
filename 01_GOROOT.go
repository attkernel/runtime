package main

import (
	"fmt"
	"runtime"
)

//GOROOT返回Go的根目录。如果存在GOROOT环境变量，返回该变量的值；否则，返回创建Go时的根目录
func main() {
	fmt.Println(runtime.GOROOT())
}

// /usr/local/Cellar/go/1.13/libexec
