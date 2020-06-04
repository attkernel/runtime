package main

import (
	"fmt"
	"runtime"
	"time"
)

type S struct {
	Name string
}

//注意x必须是指针类型,f 函数的参数一定要和x保持一致,或者写interface{},不然程序会报错
func main() {
	s := &S{}
	//runtime.SetFinalizer(s, func(i interface{}) {
	runtime.SetFinalizer(s, func(s *S) {
		fmt.Println("s被回收了")
	})
	runtime.GC()
	time.Sleep(time.Second)
}

// s被回收了
