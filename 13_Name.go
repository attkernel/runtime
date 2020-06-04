package main

import "runtime"

//获取调用栈所调用的函数的名字
func main() {
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	for _, pc := range pcs[:i] {
		funcPC := runtime.FuncForPC(pc)
		println(funcPC.Name())
	}
}

/*
main.main
runtime.main
runtime.goexit
*/
