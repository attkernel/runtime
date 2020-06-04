package main

/*
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"runtime"
)

//获取当前进程调用c方法的次数
func main() {
	fmt.Println(runtime.NumCgoCall())
}

// 1
