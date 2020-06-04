package main

import (
	"fmt"
	"runtime"
	"time"
)

type S struct {
	Name string
}

//GC执行一次垃圾回收
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
