package main

import (
	"fmt"
	"runtime"
)

//获取调用栈所调用的函数的所在的源文件名和行号
func main() {
	pcs := make([]uintptr, 10)
	nums := runtime.Callers(1, pcs)
	for i := 0; i < nums; i++ {
		pc := runtime.FuncForPC(pcs[i])
		file, line := pc.FileLine(pcs[i])
		fmt.Printf("%s:%d\n", file, line)
	}
}

/*
/Volumes/MySpace/go/src/github.com/attkernel/14_FileLine.go:10
/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:212
/usr/local/Cellar/go/1.13/libexec/src/runtime/asm_amd64.s:1358
*/
