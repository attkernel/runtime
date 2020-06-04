package main

import (
	"fmt"
	"runtime"
)

//在程序调用出执行断点并打印出相应的堆栈信息
func main() {
	fmt.Println(">>> start func")
	runtime.Breakpoint()
	fmt.Println("<<< end func")
}

/*
	>>> start func
	SIGTRAP: trace trap
	PC=0x1051071 m=0 sigcode=1

	goroutine 1 [running]:
	runtime.breakpoint()
		/usr/local/Cellar/go/1.13/libexec/src/runtime/asm_amd64.s:240 +0x1 fp=0xc00007cef0 sp=0xc00007cee8 pc=0x1051071
	runtime.Breakpoint(...)
		/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:3500
	main.main()
		/Volumes/MySpace/go/src/github.com/attkernel/08_Breakpoint.go:10 +0x7f fp=0xc00007cf60 sp=0xc00007cef0 pc=0x10992cf
	runtime.main()
		/usr/local/Cellar/go/1.13/libexec/src/runtime/proc.go:203 +0x21e fp=0xc00007cfe0 sp=0xc00007cf60 pc=0x102acbe
	runtime.goexit()
		/usr/local/Cellar/go/1.13/libexec/src/runtime/asm_amd64.s:1357 +0x1 fp=0xc00007cfe8 sp=0xc00007cfe0 pc=0x1053091

	rax    0x0
	rbx    0xc000046c00
	rcx    0x0
	rdx    0x0
	rdi    0xc00000e010
	rsi    0x0
	rbp    0xc00007cf50
	rsp    0xc00007cee8
	r8     0x4
	r9     0xfffff80000000001
	r10    0x0
	r11    0x246
	r12    0xffffffffffffffff
	r13    0xb
	r14    0xa
	r15    0x200
	rip    0x1051071
	rflags 0x206
	cs     0x2b
	fs     0x0
	gs     0x0
	exit status 2
*/
