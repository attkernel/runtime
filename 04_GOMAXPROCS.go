package main

import (
	"fmt"
	"runtime"
	"time"
)

//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。
//go 1.5 版本之前的 GOMAXPROCS 默认是 1
//go 1.5 版本之后的 GOMAXPROCS 默认是 Num of cpu
func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()
	time.Sleep(time.Second)
}

// 8
// 打印8后并不会退出因为只有一个p在运行死循环
