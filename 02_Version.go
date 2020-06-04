package main

import (
	"fmt"
	"runtime"
)

//返回Go的版本字符串。它要么是递交的hash和创建时的日期；要么是发行标签如"go1.3"
func main() {
	fmt.Println(runtime.Version())
}

// go1.13
