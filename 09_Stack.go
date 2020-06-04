package main

import (
	"fmt"
	"runtime"
	"time"
)

//Stack将调用其的go程的调用栈踪迹格式化后写入到buf中并返回写入的字节数。若all为true，函数会在写入当前go程的踪迹信息后，将其它所有go程的调用栈踪迹都格式化写入到buf中。
func main() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		select {
		case <-ticker.C:
		}
	}()
	time.Sleep(time.Second)
	buf := make([]byte, 10000)
	runtime.Stack(buf, true)
	fmt.Println(string(buf))
}

/*
--true
goroutine 1 [running]:
main.main()
	/Volumes/MySpace/go/src/github.com/attkernel/09_Stack.go:19 +0x92

goroutine 6 [chan receive]:
main.main.func1()
	/Volumes/MySpace/go/src/github.com/attkernel/09_Stack.go:14 +0x82
created by main.main
	/Volumes/MySpace/go/src/github.com/attkernel/09_Stack.go:10 +0x39

--false
goroutine 1 [running]:
main.main()
	/Volumes/MySpace/go/src/github.com/attkernel/09_Stack.go:19 +0x92
*/
