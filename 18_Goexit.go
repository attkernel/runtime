package main

import (
	"fmt"
	"runtime"
	"time"
)

//终止掉当前的go协程 Goexit终止调用它的go协程,其他协程不受影响,Goexit会在终止该go协程前执行所有的defer函数，前提是defer必须在它前面定义,如果在main go协程调用本方法,会终止该go协程,但不会让main返回,因为main函数没有返回,程序会继续执行其他go协程,当其他go协程执行完毕后,程序就会崩溃

/*func main() {
	defer func() {
		fmt.Println("defer before")
	}()
	fmt.Println(">>> main")
	runtime.Goexit()
	defer func() {
		fmt.Println("defer after")
	}()
	fmt.Println("<<< main")
}*/

func main() {
	defer func() {
		fmt.Println("defer before")
	}()
	fmt.Println(">>> main")
	go func() {
		runtime.Goexit()
	}()
	time.Sleep(time.Second)
	defer func() {
		fmt.Println("defer after")
	}()
	fmt.Println("<<< main")
}

/*
---first
>>> main
defer before
fatal error: no goroutines (main called runtime.Goexit) - deadlock!

runtime stack:
runtime.throw(0x10d6a26, 0x36)
	/usr/local/Cellar/go/1.13/libexec/src/runtime/panic.go:774 +0x72
runtime.checkdead()
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:4241 +0x30a
runtime.mput(...)
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:4627
runtime.stopm()
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:1926 +0x95
runtime.findrunnable(0xc000028f00, 0x0)
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:2391 +0x53f
runtime.schedule()
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:2524 +0x2be
runtime.park_m(0xc000000c00)
	/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:2610 +0x9d
runtime.mcall(0x80000)
	/usr/local/Cellar/go/1.13/libexec/src/runtime/asm_amd64.s:318 +0x5b
exit status 2

--second
>>> main
<<< main
defer after
defer before
*/
