/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"os"
	"time"
)

func deferClose() {
	for {
		f, _ := os.Open("README.md")
		_ = f.Close()
		defer time.Sleep(10)
	}
}
func ExampleDefer() {
	time.Sleep(time.Second * 5)
	deferClose()
}
