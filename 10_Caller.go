package main

import (
	"fmt"
	"runtime"
)

//获取当前函数或者上层函数的标识号、文件名、调用方法在当前文件中的行号 skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
func main() {
	for i := 0; i < 4; i++ {
		test(i)
	}
}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}

/*
17411743   /Volumes/MySpace/go/src/github.com/attkernel/10_Caller.go   20   true   main.call
17411647   /Volumes/MySpace/go/src/github.com/attkernel/10_Caller.go   16   true   main.test
17411638   /Volumes/MySpace/go/src/github.com/attkernel/10_Caller.go   11   true   main.main
16952509   /usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go   203   true   runtime.main
*/
