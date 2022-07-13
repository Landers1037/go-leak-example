/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

/*
#include <stdio.h>
#include <stdlib.h>
void hello(const char *s) {
	puts(s);
}
*/
import "C"
import (
	"strconv"
	"time"
)

func ExampleCgo() {
	for i := range [100000]int{} {
		str := C.CString("index:" + strconv.Itoa(i))
		C.hello(str)
		//C.free(unsafe.Pointer(str))
		time.Sleep(1 * time.Millisecond)
	}
}
