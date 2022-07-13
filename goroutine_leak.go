/*
Create: 2022/7/14
Project: go-leak-example
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	"context"
	"fmt"
	"time"
)

// goroutine泄漏

// 生产者
func produce(ch chan int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	go func() {
		ch <- 1
		fmt.Println("produce")
		time.Sleep(100)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout")
		cancel()
	}
}

func consume(ch chan int) {
	go func() {
		res := <-ch
		fmt.Println("consume", res)
	}()
}

func ExampleGoroutine() {
	time.Sleep(5 * time.Second)
	var ch = make(chan int)
	go func() {
		for {
			produce(ch)
		}
	}()

	consume(ch)
}
