/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"fmt"
	"time"
)

// 数组的内存泄漏

var s []int

// 从数组中获取一段slice 因为共用底层内存导致array无法释放
func exampleSlice(sl []int) {
	//s = sl[len(sl)-10:]
	s = append([]int(nil), sl[len(sl)-10:]...)
	fmt.Println(s)
}

func ExampleSlice() {
	var testArray = [10000000]int{}
	for i, _ := range testArray {
		testArray[i] = 10086
	}
	time.Sleep(5 * time.Second)
	exampleSlice(testArray[:])
}
