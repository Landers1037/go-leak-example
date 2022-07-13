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
	"io/ioutil"
	"net/http"
	"time"
)

func readBody() {
	for _ = range [1000]int{} {
		res, err := http.Get("http://api.renj.io")
		if err != nil {
			fmt.Println(err)
			return
		}
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(len(data))
		time.Sleep(10 * time.Millisecond)
	}
}

func ExampleHTTP() {
	readBody()
}
