package main

import (
	"fmt"
	"runtime"
)

//获取该调用栈的调用栈标识符
func main() {
	pcs := make([]uintptr, 10)
	nums := runtime.Callers(1, pcs)
	for i := 0; i < nums; i++ {
		pc := runtime.FuncForPC(pcs[i])
		res := pc.Entry()
		fmt.Printf("%d\n", res)
	}
}

/*
17411008
16952096
17117440
*/
