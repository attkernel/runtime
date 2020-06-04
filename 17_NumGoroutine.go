package main

import (
	"fmt"
	"runtime"
)

//获取当前存在的go协程数
func main() {
	go func() {}()
	go func() {}()
	fmt.Println(runtime.NumGoroutine())
}

//3
