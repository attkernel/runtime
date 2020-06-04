package main

import "runtime"

//获取函数标志符对应的函数地址
func main() {
	pcs := make([]uintptr, 10)
	i := runtime.Callers(1, pcs)
	for _, pc := range pcs[:i] {
		println(runtime.FuncForPC(pc))
	}
}

/*
0x10bf158
0x109f350
0x10b84e0
*/
