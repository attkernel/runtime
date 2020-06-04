package main

import (
	"fmt"
	"runtime"
	"time"
)

//将调用的go协程绑定到当前所在的操作系统线程，其它go协程不能进入该线程 除非调用的go程退出或调用UnlockOSThread，否则它将总是在该线程中执行，而其它go程则不能进入该线程
func main() {
	go func() {
		runtime.LockOSThread()
		fmt.Println("func 1")
		runtime.UnlockOSThread()
	}()
	go func() {
		runtime.LockOSThread()
		fmt.Println("func 2")
		runtime.UnlockOSThread()
	}()
	runtime.Gosched()
	time.Sleep(10 * time.Second)
	fmt.Println("exit")
}

/*
func 1
func 2
exit
*/
