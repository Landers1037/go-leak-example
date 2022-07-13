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

var ss string

func exampleString(str string) {
	ss = str[len(str)-5:]
	//ss = string([]byte(str[len(str)-5:]))
}

func ExampleString() {
	var testStr = "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"
	time.Sleep(5 * time.Second)
	exampleString(testStr)
}
