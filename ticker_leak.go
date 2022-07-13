/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"time"
)

// ticker泄漏

func tickerRun() {
	for _ = range [100000]int{} {
		ticker := time.NewTicker(1 * time.Millisecond)
		defer ticker.Stop()
		//fmt.Println(ticker)
	}
}

func ExampleTicker() {
	tickerRun()
}
