package main

import (
	"fmt"
	"runtime"
)

//NumCPU返回本地机器的逻辑CPU个数
func main() {
	fmt.Println(runtime.NumCPU())
}

// 8
