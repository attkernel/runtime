package main

import (
	"fmt"
	"runtime"
)

//Callers用来返回调用站的程序计数器, 放到一个uintptr中。0 代表 Callers 本身，这和Caller的参数的意义不一样，历史原因造成的。 1 才对应Caller的 0。
func main() {
	test(0)
}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	res := make([]uintptr, 10)
	nums := runtime.Callers(skip, res)
	for i := 0; i < nums; i++ {
		f := runtime.FuncForPC(res[i])
		file, line := f.FileLine(res[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

/*func main() {
	res := make([]uintptr, 10)
	nums := runtime.Callers(0, res)
	for i := 0; i < nums; i++ {
		f := runtime.FuncForPC(res[i])
		file, line := f.FileLine(res[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
		//fmt.Printf("name: %s,pointor: %p\n", runtime.FuncForPC(v).Name(), v)
	}
}*/

/*
/usr/local/Cellar/go/1.13/libexec/src/runtime/extern.go:211 runtime.Callers
/Volumes/MySpace/go/src/github.com/attkernel/11_Callers.go:19 main.call
/Volumes/MySpace/go/src/github.com/attkernel/11_Callers.go:14 main.main
/Volumes/MySpace/go/src/github.com/attkernel/11_Callers.go:10 main.main
/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:212 runtime.main
/usr/local/Cellar/go/1.13/libexec/src/runtime/asm_amd64.s:1358 runtime.goexit
*/
