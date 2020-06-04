package main

import (
	"fmt"
	"runtime"
	"time"
)

//暂停当前goroutine，使其他goroutine先行运算。只是暂停，不是挂起，当时间片轮转到该协程时，Gosched()后面的操作将自动恢复 主协程不是一定要等其他协程执行完才会继续执行，而是一定时间。如果这个时间内其他协程没有执行完，那么主协程将继续执行
//---first
/*func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	go say("world")
	say("hello")
}*/
//--second
func say(s string) {
	time.Sleep(time.Minute)
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go say("func")
	runtime.Gosched()
	fmt.Println("exit")
}

/*
--first
hello
world
hello
--second
exit
*/
